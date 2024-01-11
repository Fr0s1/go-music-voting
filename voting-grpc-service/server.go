package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"

	pb "voting-grpc/pkg/grpc"

	"google.golang.org/grpc"

	database "voting-grpc/pkg/db/mysql"

	grpc_user_client "voting-grpc/pkg/grpc/clients/user-service"

	model "voting-grpc/pkg/db/models"

	logging "voting-grpc/pkg/logging"

	"google.golang.org/grpc/credentials"
)

var (
	port   = flag.Int("port", 50052, "The server port")
	logger = logging.Log.WithFields(logging.StandardFields)
)

type VotingServer struct {
	pb.UnimplementedVotingServer

	mu         sync.Mutex
	PollAlbums map[int64][]*pb.NewPollAlbumResponse
}

func (s *VotingServer) CreatePoll(ctx context.Context, in *pb.NewPoll) (*pb.Poll, error) {
	creator_id, poll_name := in.CreatorId, in.Name

	grpc_ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	user, _ := grpc_user_client.GrpcClient.GetUserDetails(grpc_ctx, &pb.UserQuery{UserId: creator_id})

	defer cancel()

	var poll_albums []*model.Album

	poll := &model.Poll{
		Name: poll_name,
		Creator: &model.User{
			Id:       user.Id,
			Username: user.Username,
		},
		Albums: poll_albums,
	}

	poll_id := poll.Save()

	var grpc_poll_albums []*pb.Album

	return &pb.Poll{
		Name:      poll.Name,
		Id:        poll_id,
		Albums:    grpc_poll_albums,
		CreatorId: creator_id,
	}, nil
}

func (s *VotingServer) AddPollAlbums(stream pb.Voting_AddPollAlbumsServer) error {
	for {
		pollAlbum, err := stream.Recv()

		logger.Info("AddPollAlbums gRPC Server: received albums from server", pollAlbum)

		if pollAlbum == nil {
			return nil
		}
		pollId := pollAlbum.PollId
		albumId := pollAlbum.AlbumId

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		_ = model.AddAlbumPoll(albumId, pollId)

		s.mu.Lock()
		albumModel := model.GetAlbumDetails(albumId)

		newPollAlbumResponse := &pb.NewPollAlbumResponse{
			PollId: pollId,
			Album: &pb.Album{
				Id:     albumModel.Id,
				Name:   albumModel.Name,
				Artist: albumModel.Artist,
			},
		}

		s.PollAlbums[pollId] = append(s.PollAlbums[pollId], newPollAlbumResponse)

		rn := make([]*pb.NewPollAlbumResponse, len(s.PollAlbums[pollId]))
		copy(rn, s.PollAlbums[pollId])
		s.mu.Unlock()

		for _, albumResponse := range rn {
			if err := stream.Send(albumResponse); err != nil {
				return err
			}
		}
	}
}

func (s *VotingServer) VoteAlbum(ctx context.Context, in *pb.Vote) (*pb.Vote, error) {
	_, err := model.VotePollAlbum(in.PollId, in.AlbumId, in.UserId)

	if err != nil {
		logger.Error(err)

		return &pb.Vote{}, nil
	}

	return in, nil
}

func (s *VotingServer) GetPollDetails(ctx context.Context, in *pb.PollQuery) (*pb.PollDetails, error) {
	pollId := in.PollId

	logger.Info("GetPollDetails: Reach here")
	rows, err := database.Db.Query("SELECT p.Name as PollName, p.CreatorID, a.ID, a.Artist, a.Name FROM Polls p LEFT JOIN Poll_Album pa on p.ID = pa.PollID LEFT JOIN Albums a on pa.AlbumID = a.ID WHERE p.ID= ?", pollId)

	// Keep track
	totalAlbum := 0

	if err != nil {
		logger.Error(err)
	}

	defer rows.Close()
	logger.Info("GetPollDetails: Reach here 2")

	var pollName string
	var creatorId int32

	albumVotesChan := make(chan map[model.Album][]*pb.Vote)

	wg := sync.WaitGroup{}

	for rows.Next() {
		var album model.Album

		if err := rows.Scan(&pollName, &creatorId, &album.Id, &album.Artist, &album.Name); err != nil {
			logger.Error(err)
		}

		logger.Info("GetPollDetails: Reach here 3")
		logger.Info("Album value: %v", album)

		wg.Add(1)
		ctx, cancelFunc := context.WithTimeoutCause(ctx, time.Second*1, errors.New("query db takes too long"))
		defer cancelFunc()

		go func(pollId int64, album model.Album, albumVotesChan chan<- map[model.Album][]*pb.Vote, ctx context.Context, wg *sync.WaitGroup) {
			defer wg.Done()

			// time.Sleep(time.Second * 2)
			s.mu.Lock()
			totalAlbum += 1
			s.mu.Unlock()

			select {
			case <-ctx.Done():
				logger.Info("Error when querying database: ", ctx.Err())
				return
			default:
				logger.Info("GetPollDetails Routine: Reach here 1")
				rowsRoutine, _ := database.Db.Query("SELECT VoterID, PollID, AlbumID FROM Votes WHERE PollID = ? and AlbumID = ?", pollId, album.Id)
				defer rowsRoutine.Close()
				logger.Info("GetPollDetails Routine: Reach here 2")

				var vote pb.Vote

				albumVotesMap := make(map[model.Album][]*pb.Vote)

				albumVotesMap[album] = []*pb.Vote{}

				for rowsRoutine.Next() {
					logger.Info("GetPollDetails Routine: Reach here 3")
					if err := rowsRoutine.Scan(&vote.UserId, &vote.PollId, &vote.AlbumId); err != nil {
						logger.Error(err)

						continue
					}
					logger.Info("GetPollDetails Routine: Reach here 4")

					fmt.Printf("Value: %+v\n", vote)

					albumVotesMap[album] = append(albumVotesMap[album], &vote)
				}
				logger.Info("GetPollDetails Routine: Reach here 5")
				albumVotesChan <- albumVotesMap
				logger.Info("GetPollDetails Routine: Reach here 6")
			}

		}(pollId, album, albumVotesChan, ctx, &wg)
	}
	logger.Info("GetPollDetails: Reach here 4")
	// wg.Wait()
	// close(albumVotesChan)
	// logger.Info("GetPollDetails: Reach here 6")

	poll := &pb.PollDetails{}

	poll.Name = pollName
	poll.Id = pollId
	poll.CreatorId = creatorId

	var albumsVotesInPoll []*pb.PollDetails_AlbumVote

	logger.Info("GetPollDetails: Reach here 7")
	logger.Info("Total album: ", totalAlbum)

	go func() {
		wg.Wait()
		logger.Info("GetPollDetails: Reach here 5")
		close(albumVotesChan)
	}()

	for i := 0; i < totalAlbum; i++ {
		albumVotes := <-albumVotesChan
		logger.Info("GetPollDetails: Reach here 8")
		logger.Info("Album votes: %v", albumVotes)

		pollAlbumVotes := &pb.PollDetails_AlbumVote{}

		for album, votes := range albumVotes {
			logger.Info("Album value: %v", album)

			pollAlbumVotes.Album = &pb.Album{
				Id:     album.Id,
				Name:   album.Name,
				Artist: album.Artist,
			}

			pollAlbumVotes.Votes = votes
		}

		if pollAlbumVotes.Album != nil {
			albumsVotesInPoll = append(albumsVotesInPoll, pollAlbumVotes)
		}
	}

	poll.AlbumVotes = albumsVotesInPoll

	return poll, nil
}

func main() {
	flag.Parse()

	currentWorkDir, _ := os.Getwd()
	cert, err := tls.LoadX509KeyPair(filepath.Join(currentWorkDir, "pkg/tls/voting-grpc-cert.pem"), filepath.Join(currentWorkDir, "./pkg/tls/voting-grpc-key.pem"))

	if err != nil {
		logger.Fatal("failed to load key pair: %s", err)
	}

	ca := x509.NewCertPool()
	caFilePath := filepath.Join(currentWorkDir, "./pkg/tls/ca-cert.pem")
	caBytes, err := os.ReadFile(caFilePath)

	if err != nil {
		logger.Fatal("failed to read ca cert %q: %v", caFilePath, err)
	}

	if ok := ca.AppendCertsFromPEM(caBytes); !ok {
		logger.Fatal("failed to parse ", &caFilePath)
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.NoClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    ca,
	}

	grpcServer := grpc.NewServer(grpc.Creds(credentials.NewTLS(tlsConfig)))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database.InitDB()

	defer database.CloseDB()

	grpc_user_client.InitConnection()

	s := &VotingServer{}

	s.PollAlbums = make(map[int64][]*pb.NewPollAlbumResponse)

	pb.RegisterVotingServer(grpcServer, s)

	grpcServer.Serve(lis)
}

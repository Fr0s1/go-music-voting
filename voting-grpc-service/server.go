package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"

	pb "voting-grpc/pkg/grpc"

	"google.golang.org/grpc"

	database "voting-grpc/pkg/db/mysql"

	grpc_user_client "voting-grpc/pkg/grpc/clients/user-service"

	model "voting-grpc/pkg/db/models"
)

var (
	port = flag.Int("port", 50052, "The server port")
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
		log.Fatal(err)
	}

	return in, nil
}

func (s *VotingServer) GetPollDetails(ctx context.Context, in *pb.PollQuery) (*pb.PollDetails, error) {
	pollId := in.PollId

	fmt.Println("GetPollDetails: Reach here")
	rows, err := database.Db.Query("SELECT p.Name, p.CreatorID, a.ID, a.Artist, a.Name FROM Polls p JOIN Poll_Album pa on p.ID = pa.PollID JOIN Albums a on pa.AlbumID = a.ID WHERE p.ID= ?", pollId)

	// Keep track
	totalAlbum := 0

	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	fmt.Println("GetPollDetails: Reach here 2")

	var pollName string
	var creatorId int32

	albumVotesChan := make(chan map[model.Album][]*pb.Vote)

	// wg := sync.WaitGroup{}

	for rows.Next() {
		var album model.Album
		totalAlbum += 1
		if err := rows.Scan(&pollName, &creatorId, &album.Id, &album.Artist, &album.Name); err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("GetPollDetails: Reach here 3")
		fmt.Println("Album value: %v", album)

		// wg.Add(1)
		ctx, cancelFunc := context.WithTimeoutCause(ctx, time.Second*2, errors.New("query db takes too long"))
		defer cancelFunc()

		go func(pollId int64, albumId int64, albumVotesChan chan<- map[model.Album][]*pb.Vote, ctx context.Context) {
			select {
			case <-ctx.Done():
				fmt.Println("Error when querying database: ", ctx.Err())

				return
			default:
				fmt.Println("GetPollDetails Routine: Reach here 1")
				rows_routine, _ := database.Db.Query("SELECT AlbumID, PollID, VoterID FROM Votes WHERE PollID = ? and AlbumID = ?", pollId, albumId)
				defer rows_routine.Close()
				fmt.Println("GetPollDetails Routine: Reach here 2")

				var vote pb.Vote

				albumVotesMap := make(map[model.Album][]*pb.Vote)

				for rows_routine.Next() {
					fmt.Println("GetPollDetails Routine: Reach here 3")
					if err := rows_routine.Scan(&vote.UserId, &vote.PollId, &vote.AlbumId); err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println("GetPollDetails Routine: Reach here 4")

					fmt.Printf("Value: %+v\n", vote)

					albumVotesMap[album] = append(albumVotesMap[album], &vote)
				}
				fmt.Println("GetPollDetails Routine: Reach here 5")
				albumVotesChan <- albumVotesMap
				fmt.Println("GetPollDetails Routine: Reach here 6")
			}
			// defer wg.Done()

		}(pollId, album.Id, albumVotesChan, ctx)
	}
	fmt.Println("GetPollDetails: Reach here 4")
	// wg.Wait()
	// fmt.Println("GetPollDetails: Reach here 5")
	// close(albumVotesChan)
	// fmt.Println("GetPollDetails: Reach here 6")

	poll := &pb.PollDetails{}

	poll.Name = pollName
	poll.Id = pollId
	poll.CreatorId = creatorId

	var albumsVotesInPoll []*pb.PollDetails_AlbumVote

	fmt.Println("GetPollDetails: Reach here 7")

	for i := 0; i < totalAlbum; i++ {
		albumVotes := <-albumVotesChan
		fmt.Println("GetPollDetails: Reach here 8")

		pollAlbumVotes := &pb.PollDetails_AlbumVote{}

		for album, votes := range albumVotes {
			pollAlbumVotes.Album = &pb.Album{
				Id:     album.Id,
				Name:   album.Name,
				Artist: album.Artist,
			}

			pollAlbumVotes.Votes = votes
		}

		albumsVotesInPoll = append(albumsVotesInPoll, pollAlbumVotes)
	}

	poll.AlbumVotes = albumsVotesInPoll

	return poll, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database.InitDB()

	defer database.CloseDB()

	grpc_user_client.InitConnection()

	grpcServer := grpc.NewServer()

	s := &VotingServer{}

	s.PollAlbums = make(map[int64][]*pb.NewPollAlbumResponse)

	pb.RegisterVotingServer(grpcServer, s)

	grpcServer.Serve(lis)
}

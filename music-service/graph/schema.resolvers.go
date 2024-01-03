package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"music-service/graph/model"
	"music-service/pkg/albums"
	"music-service/pkg/auth"
	pb "music-service/pkg/grpc"
	grpc_client "music-service/pkg/grpc/client"
	"music-service/pkg/logging"
	"music-service/pkg/users"
	"strconv"
	"time"
)

// UploadAlbum is the resolver for the uploadAlbum field.
func (r *mutationResolver) UploadAlbum(ctx context.Context, input model.NewAlbum) (*model.Album, error) {
	user := auth.ForContext(ctx)

	logging.Log.WithFields(logging.StandardFields).Info("User: ", user)

	user_json, _ := json.Marshal(user)

	logging.Log.WithFields(logging.StandardFields).Info("User info: ", string(user_json))

	logging.Log.WithFields(logging.StandardFields).Info("Reach UploadAlbum function")

	if user == nil {
		return &model.Album{}, fmt.Errorf("access denied")
	}

	var album albums.Album

	album.Name = input.Name
	album.Artist = input.Artist
	album.Genre = input.Genre
	album.Year = input.Year

	album.Uploader = user

	logging.Log.WithFields(logging.StandardFields).Info("Reach insert statement: album ", album)

	albumId := album.Save()

	return &model.Album{ID: strconv.Itoa(int(albumId)), Name: album.Name, Artist: album.Artist, Genre: album.Genre, Year: album.Year, Uploader: &model.User{ID: user.Id, Username: user.Username}}, nil
}

// CreatePoll is the resolver for the createPoll field.
func (r *mutationResolver) CreatePoll(ctx context.Context, input model.NewPoll) (*model.Poll, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return &model.Poll{}, fmt.Errorf("access denied")
	}

	userId, _ := strconv.ParseInt(user.Id, 10, 32)

	newUserId := int32(userId)

	albumsList := input.Albums

	grpcNewPoll := &pb.NewPoll{
		Name:      input.Name,
		CreatorId: newUserId,
	}

	grpcCtx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	grpcPoll, err := grpc_client.VotingGRPCClient.CreatePoll(grpcCtx, grpcNewPoll)

	if err != nil {
		log.Printf(err.Error())
	}

	stream, err := grpc_client.VotingGRPCClient.AddPollAlbums(ctx)

	if err != nil {
		log.Printf(err.Error())
	}

	// Slice to save the list of albums returned by voting gRPC server
	grpcCurrentPollAlbums := []*pb.NewPollAlbumResponse{}

	if len(albumsList) > 0 {
		waitc := make(chan struct{})
		go func() {
			for {
				in, err := stream.Recv()

				if err == io.EOF {
					// read done.
					close(waitc)
					return
				}

				log.Printf("Current added albums: %v\n", in)

				grpcCurrentPollAlbums = append(grpcCurrentPollAlbums, in)
			}
		}()

		for _, album := range albumsList {
			graphqlAlbum := &albums.Album{
				Name:   album.Name,
				Artist: album.Artist,
				Genre:  album.Genre,
				Year:   album.Year,
				Uploader: &users.User{
					Id:       user.Id,
					Username: user.Username,
				},
			}

			fmt.Printf("CreatePoll GraphQL: graphqlAlbum: %v\n", graphqlAlbum)

			albumId := graphqlAlbum.Save()

			fmt.Printf("CreatePoll GraphQL: added Album ID: %v\n", graphqlAlbum)

			if albumId > 0 {
				newProtoBufAlbum := &pb.NewPollAlbum{
					PollId:  grpcPoll.Id,
					AlbumId: albumId,
				}

				stream.Send(newProtoBufAlbum)
			}
		}

		fmt.Println("CreatePoll GraphQL: Reach before close channel statement")
		stream.CloseSend()
		fmt.Println("CreatePoll GraphQL: Reach after close channel statement")

		<-waitc

		fmt.Printf("List of added albums: %v", grpcCurrentPollAlbums)
	}

	return &model.Poll{
		ID:         strconv.Itoa(int(grpcPoll.Id)),
		Name:       grpcNewPoll.Name,
		AlbumVotes: []*model.PollAlbum{},
	}, nil
}

// VoteAlbum is the resolver for the voteAlbum field.
func (r *mutationResolver) VoteAlbum(ctx context.Context, input model.NewVote) (*model.Vote, error) {
	user := auth.ForContext(ctx)

	if user == nil {
		return &model.Vote{}, fmt.Errorf("access denied")
	}

	pollId, _ := strconv.ParseInt(input.PollID, 10, 64)
	albumId, _ := strconv.ParseInt(input.AlbumID, 10, 64)
	userId, _ := strconv.ParseInt(user.Id, 10, 32)

	newUserId := int32(userId)

	pbVote := &pb.Vote{
		PollId:  pollId,
		AlbumId: albumId,
		UserId:  newUserId,
	}

	_, err := grpc_client.VotingGRPCClient.VoteAlbum(ctx, pbVote)

	if err != nil {
		return &model.Vote{}, err
	}

	graphVote := &model.Vote{
		PollID:  strconv.Itoa(int(pollId)),
		AlbumID: strconv.Itoa(int(albumId)),
	}

	return graphVote, nil
}

// GetAlbum is the resolver for the getAlbum field.
func (r *queryResolver) GetAlbum(ctx context.Context, input model.AlbumSearch) ([]*model.Album, error) {
	album, err := albums.GetAlbum(input.Name, input.Artist)

	fmt.Println("Reach GetAlbum resolver")

	var albums []*model.Album

	if err != nil {
		log.Fatal(err)

		return albums, nil
	}

	if album.ID == "" {
		return albums, nil
	}

	graphql_album := &model.Album{ID: album.ID, Name: album.Name, Artist: album.Artist, Genre: album.Genre, Year: album.Year, Uploader: &model.User{ID: album.Uploader.Id, Username: album.Uploader.Username}}

	albums = append(albums, graphql_album)

	return albums, nil
}

// GetAllAlbums is the resolver for the getAllAlbums field.
func (r *queryResolver) GetAllAlbums(ctx context.Context) ([]*model.Album, error) {
	var resultAlbums []*model.Album

	albums := albums.GetAll()

	for _, album := range albums {
		uploader := &model.User{
			ID:       album.Uploader.Id,
			Username: album.Uploader.Username,
		}

		graphql_album := &model.Album{
			ID:       album.ID,
			Name:     album.Name,
			Year:     album.Year,
			Artist:   album.Artist,
			Genre:    album.Genre,
			Uploader: uploader,
		}

		resultAlbums = append(resultAlbums, graphql_album)
	}

	return resultAlbums, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

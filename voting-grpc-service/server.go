package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
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

	pb.RegisterVotingServer(grpcServer, s)

	grpcServer.Serve(lis)
}

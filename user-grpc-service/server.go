package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "user-grpc/pkg/grpc"
	jwt "user-grpc/pkg/jwt"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type UserCredentialsServer struct {
	pb.UnimplementedUserCredentialsServer
}

func (s *UserCredentialsServer) GetUser(ctx context.Context, in *pb.UserJWTToken) (*pb.User, error) {

	user, _ := jwt.ParseToken(in.Token)

	return &user, nil
}

func main() {
	fmt.Print("test")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := &UserCredentialsServer{}

	pb.RegisterUserCredentialsServer(grpcServer, s)

	grpcServer.Serve(lis)
}

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

	database "user-grpc/pkg/db/mysql"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type UserCredentialsServer struct {
	pb.UnimplementedUserCredentialsServer
}

func (s *UserCredentialsServer) GetUser(ctx context.Context, in *pb.UserJWTToken) (*pb.User, error) {

	user, err := jwt.ParseToken(in.Token)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserCredentialsServer) GetUserDetails(ctx context.Context, in *pb.UserQuery) (*pb.User, error) {
	user_id := in.UserId

	stmt, err := database.Db.Prepare("SELECT Username FROM Users WHERE ID = (?)")

	if err != nil {
		log.Fatal(err)
	}

	row := stmt.QueryRow(user_id)

	if err != nil {
		log.Fatal(err)
	}

	var username string

	err = row.Scan(&username)

	return &pb.User{Id: user_id, Username: username}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	database.InitDB()

	defer database.CloseDB()

	grpcServer := grpc.NewServer()

	s := &UserCredentialsServer{}

	pb.RegisterUserCredentialsServer(grpcServer, s)

	grpcServer.Serve(lis)
}

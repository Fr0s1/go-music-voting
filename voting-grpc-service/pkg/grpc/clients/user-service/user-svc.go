package grpc

import (
	"flag"
	"log"

	pb "voting-grpc/pkg/grpc"

	"google.golang.org/grpc"

	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var GrpcConnection any

var GrpcClient pb.UserCredentialsClient

func InitConnection() {
	flag.Parse()

	// Set up connection to gRPC user credential service
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	GrpcConnection = conn

	c := pb.NewUserCredentialsClient(conn)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	GrpcClient = c
}

// func CloseConnection() {
// 	GrpcConnection.Close()
// }

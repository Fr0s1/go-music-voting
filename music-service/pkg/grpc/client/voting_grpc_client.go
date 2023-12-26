package client

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	pb "music-service/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	votingGRPCAddr = flag.String("votingaddr", "localhost:50052", "the address to connect to")
)

var VotingGRPCConnection any

var VotingGRPCClient pb.VotingClient

func InitVotingConnection() {
	flag.Parse()
	currentWorkDir, _ := os.Getwd()

	cert, err := tls.LoadX509KeyPair(filepath.Join(currentWorkDir, "pkg/tls/music-service-cert.pem"), filepath.Join(currentWorkDir, "pkg/tls/music-service-key.pem"))

	if err != nil {
		log.Fatalf("failed to load client cert: %v", err)
	}

	ca := x509.NewCertPool()
	caFilePath := filepath.Join(currentWorkDir, "pkg/tls/ca-cert.pem")
	caBytes, err := os.ReadFile(caFilePath)

	if err != nil {
		log.Fatalf("failed to read ca cert %q: %v", caFilePath, err)
	}
	if ok := ca.AppendCertsFromPEM(caBytes); !ok {
		log.Fatalf("failed to parse %q", caFilePath)
	}

	tlsConfig := &tls.Config{
		ServerName:   "voting.grpc.cloudfrosted.com",
		Certificates: []tls.Certificate{cert},
		RootCAs:      ca,
	}

	fmt.Println("Start connecting to voting gRPC service addr ", *votingGRPCAddr)

	// Set up connection to gRPC voting credential service
	conn, conn_err := grpc.Dial(*votingGRPCAddr, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))

	fmt.Printf("Error: %v\n", conn_err)
	VotingGRPCConnection = conn

	c := pb.NewVotingClient(conn)

	if conn_err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	VotingGRPCClient = c
}

// func CloseConnection() {
// 	GrpcConnection.Close()
// }

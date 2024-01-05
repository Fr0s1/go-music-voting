package client

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"os"
	"path/filepath"

	pb "voting-grpc/pkg/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	logging "voting-grpc/pkg/logging"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

var GrpcConnection any

var GrpcClient pb.UserCredentialsClient

func InitConnection() {
	logger := logging.Log.WithFields(logging.StandardFields)

	flag.Parse()
	currentWorkDir, _ := os.Getwd()

	cert, err := tls.LoadX509KeyPair(filepath.Join(currentWorkDir, "pkg/tls/voting-grpc-cert.pem"), filepath.Join(currentWorkDir, "pkg/tls/voting-grpc-key.pem"))

	if err != nil {
		logger.Error("failed to load client cert: %v", err)
	}

	ca := x509.NewCertPool()
	caFilePath := filepath.Join(currentWorkDir, "pkg/tls/ca-cert.pem")
	caBytes, err := os.ReadFile(caFilePath)

	if err != nil {
		logger.Error("failed to read ca cert %q: %v", caFilePath, err)
	}
	if ok := ca.AppendCertsFromPEM(caBytes); !ok {
		logger.Error("failed to parse %q", caFilePath)
	}

	tlsConfig := &tls.Config{
		ServerName:   "user.grpc.cloudfrosted.com",
		Certificates: []tls.Certificate{cert},
		RootCAs:      ca,
	}

	logger.Info("Start connecting to user gRPC service addr ", *addr)

	// Set up connection to gRPC user credential service
	conn, conn_err := grpc.Dial(*addr, grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)))

	logger.Info("Error: ", conn_err)
	GrpcConnection = conn

	c := pb.NewUserCredentialsClient(conn)

	if conn_err != nil {
		logger.Error("did not connect: %v", err)
	}

	GrpcClient = c
}

// func CloseConnection() {
// 	GrpcConnection.Close()
// }

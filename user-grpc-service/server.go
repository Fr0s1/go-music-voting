package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	pb "user-grpc/pkg/grpc"
	jwt "user-grpc/pkg/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	database "user-grpc/pkg/db/mysql"

	logging "user-grpc/pkg/logging"
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

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	logging.Log.WithFields(logging.StandardFields).Info("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func main() {
	flag.Parse()

	logger := logging.Log.WithFields(logging.StandardFields)

	// cert, err := tls.LoadX509KeyPair(data.Path("./pkg/tls/server-cert.pem"), data.Path("./pkg/tls/server-cert.key"))
	currentWorkDir, _ := os.Getwd()
	cert, err := tls.LoadX509KeyPair(filepath.Join(currentWorkDir, "pkg/tls/server-cert.pem"), filepath.Join(currentWorkDir, "./pkg/tls/server-key.pem"))

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
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		ClientCAs:    ca,
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewTLS(tlsConfig)),
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	s := &UserCredentialsServer{}

	pb.RegisterUserCredentialsServer(grpcServer, s)

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost: %d", *port))

	if err != nil {
		logger.Fatal("failed to listen: ", err)
	}

	database.InitDB()

	defer database.CloseDB()

	grpcServer.Serve(lis)
}

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"music-service/graph"

	pb "music-service/pkg/grpc"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	grpc_client "music-service/pkg/grpc/client"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	grpc_client.InitConnection()

	defer grpc_client.CloseConnection()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDMwNjQyOTgsInVzZXJuYW1lIjoiZnJvc3QifQ.kpB90p7E07MbwI0zSNi2zdtY2-A4WSbABL3C1P_8zLc"

	r, _ := grpc_client.GrpcClient.GetUser(ctx, &pb.UserJWTToken{Token: token})

	fmt.Printf("Info: %+v\n", r)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

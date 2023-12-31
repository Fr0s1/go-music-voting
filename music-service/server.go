package main

import (
	"net/http"
	"os"

	"music-service/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	grpc_client "music-service/pkg/grpc/client"

	database "music-service/pkg/db/mysql"

	"music-service/pkg/auth"

	"github.com/go-chi/chi/v5"

	logging "music-service/pkg/logging"
)

const defaultPort = "8000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	router.Use(auth.Middleware())

	grpc_client.InitConnection()

	database.InitDB()

	defer database.CloseDB()

	// defer grpc_client.CloseConnection()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	logging.Log.WithFields(logging.StandardFields).Info("connect to http://localhost:%s/ for GraphQL playground", port)
	logging.Log.WithFields(logging.StandardFields).Fatal(http.ListenAndServe(":"+port, router))
}

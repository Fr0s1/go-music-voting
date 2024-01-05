package auth

import (
	"context"
	"fmt"
	grpc_client "music-service/pkg/grpc/client"
	"net/http"
	"strconv"
	"time"

	pb "music-service/pkg/grpc"
	"music-service/pkg/users"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Reach middleware here")
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			fmt.Println("Reach middleware here 1")
			grpc_ctx, cancel := context.WithTimeout(context.Background(), time.Second)

			defer cancel()

			tokenStr := header

			grpc_user, err := grpc_client.GrpcClient.GetUser(grpc_ctx, &pb.UserJWTToken{Token: tokenStr})

			fmt.Printf("gRPC User response: %v\n", grpc_user)

			var user *users.User

			if err == nil {
				user = &users.User{Id: strconv.Itoa(int(grpc_user.Id)), Username: grpc_user.Username}
			}

			fmt.Println("Reach middleware here 3")

			fmt.Printf("User: %v", user)

			graphql_ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(graphql_ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func ForContext(ctx context.Context) *users.User {
	raw, _ := ctx.Value(userCtxKey).(*users.User)

	return raw
}

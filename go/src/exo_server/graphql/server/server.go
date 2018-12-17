//go:generate go run $GOPATH/src/exo_server/scripts/gqlgen.go -v

package main

import (
	ppb "exo/project/pb"
	upb "exo/user/pb"
	"exo_server/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"google.golang.org/grpc"
)

const defaultPort = "8080"

const (
	projectPort = "50051"
	userPort    = "50052"
)

func createResolver() (*graphql.Resolver, error) {
	uconn, err := grpc.Dial(":"+userPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	uclient := upb.NewUserServiceClient(uconn)
	pconn, err := grpc.Dial(":"+projectPort, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	pclient := ppb.NewProjectServiceClient(pconn)
	return &graphql.Resolver{UserClient: uclient, ProjectClient: pclient}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	resolver, err := createResolver()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{Resolvers: resolver})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

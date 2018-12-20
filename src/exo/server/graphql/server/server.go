//go:generate go run $GOPATH/src/exo/server/scripts/gqlgen.go -v

package main

import (
	"exo/server/graphql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "8080"

const (
	projectPort = "50051"
	userPort    = "50052"
)

func createResolver() (*graphql.Resolver, error) {
	r := &graphql.Resolver{UserPort: userPort, ProjectPort: projectPort}
	if err := r.Connect(); err != nil {
		return nil, err
	}
	return r, nil
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

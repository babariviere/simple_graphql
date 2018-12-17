//go:generate protoc -I $GOPATH/src/exo/project/pb --go_out=plugins=grpc:$GOPATH/src/exo/project/pb $GOPATH/src/exo/project/pb/project.proto

package main

import (
	"exo/project"
	"exo/project/pb"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "50051"
	nbProjects  = 10
)

type server struct{}

func (s *server) AllProjects(req *pb.ProjectsRequest, result pb.ProjectService_AllProjectsServer) error {
	for i := 0; i < nbProjects; i++ {
		pj := project.Project{ID: fmt.Sprint(i), Name: fmt.Sprint("project", i)}
		if err := result.Send(pj.ToGRPC()); err != nil {
			return err
		}
	}
	return nil
}

func startServer() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}

	s := grpc.NewServer()
	pb.RegisterProjectServiceServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}
	return nil
}

func main() {
	if err := startServer(); err != nil {
		log.Fatal(err)
	}
}

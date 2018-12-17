//go:generate protoc -I $GOPATH/src/exo/user/pb --go_out=plugins=grpc:$GOPATH/src/exo/user/pb $GOPATH/src/exo/user/pb/user.proto

package main

import (
	"exo/user"
	"exo/user/pb"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "50052"
	nbUsers     = 10
)

type server struct{}

func (s *server) AllUsers(req *pb.UsersRequest, result pb.UserService_AllUsersServer) error {
	for i := 0; i < nbUsers; i++ {
		u := user.User{ID: fmt.Sprint(i), Email: fmt.Sprintf("test%d@address.com", i)}
		if err := result.Send(u.ToGRPC()); err != nil {
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
	pb.RegisterUserServiceServer(s, &server{})
	reflection.Register(s)
	fmt.Println("listening to :" + port)
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

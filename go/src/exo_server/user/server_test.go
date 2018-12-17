package main

import (
	"context"
	"exo/user"
	"exo/user/pb"
	"fmt"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
)

// return a function to close server
func startTestServer(t *testing.T) func() {
	lis, err := net.Listen("tcp", ":"+defaultPort)
	if err != nil {
		t.Errorf("failed to listen: %v", err)
		return nil
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	teardown := func() {
		s.Stop()
		lis.Close()
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			t.Errorf("failed to serve: %v", err)
		}
	}()
	return teardown
}

func newTestClient(t *testing.T) (pb.UserServiceClient, func()) {
	conn, err := grpc.Dial(":"+defaultPort, grpc.WithInsecure())
	if err != nil {
		t.Errorf("failed to dial: %v", err)
	}
	teardown := func() {
		conn.Close()
	}
	return pb.NewUserServiceClient(conn), teardown
}

func testAllUsers(t *testing.T) {
	c, teardown := newTestClient(t)
	defer teardown()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.AllUsers(ctx, &pb.UsersRequest{})
	if err != nil {
		t.Errorf("failed to get users: %v", err)
	}
	for i := 0; i < nbUsers; i++ {
		expected := user.User{ID: fmt.Sprint(i), Email: fmt.Sprintf("test%d@address.com", i)}
		p, err := r.Recv()
		if err != nil {
			t.Errorf("failed to recv user: %v", err)
		}
		grpc := user.FromGRPC(p)
		if grpc != expected {
			t.Errorf("%v != %v", grpc, expected)
		}
	}
}

func TestServer(t *testing.T) {
	teardown := startTestServer(t)

	t.Run("AllUsers", testAllUsers)

	teardown()

}

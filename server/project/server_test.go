package main

import (
	"context"
	"exo/project"
	"exo/project/pb"
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
	pb.RegisterProjectServiceServer(s, &server{})
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

func newTestClient(t *testing.T) (pb.ProjectServiceClient, func()) {
	conn, err := grpc.Dial(":"+defaultPort, grpc.WithInsecure())
	if err != nil {
		t.Errorf("failed to dial: %v", err)
	}
	teardown := func() {
		conn.Close()
	}
	return pb.NewProjectServiceClient(conn), teardown
}

func testAllProjects(t *testing.T) {
	c, teardown := newTestClient(t)
	defer teardown()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.AllProjects(ctx, &pb.ProjectsRequest{})
	if err != nil {
		t.Errorf("failed to get projects: %v", err)
	}
	for i := 0; i < nbProjects; i++ {
		expected := project.Project{ID: fmt.Sprint(i), Name: fmt.Sprint("project", i)}
		p, err := r.Recv()
		if err != nil {
			t.Errorf("failed to recv project: %v", err)
		}
		grpc := project.FromGRPC(p)
		if grpc != expected {
			t.Errorf("%v != %v", grpc, expected)
		}
	}
}

func TestServer(t *testing.T) {
	teardown := startTestServer(t)

	t.Run("AllProjects", testAllProjects)

	teardown()

}

package graphql

import (
	"context"
	"exo/project"
	ppb "exo/project/pb"
	"exo/user"
	upb "exo/user/pb"
	"io"
	"log"

	"google.golang.org/grpc"
)

type Resolver struct {
	ProjectClient ppb.ProjectServiceClient
	ProjectPort   string
	UserClient    upb.UserServiceClient
	UserPort      string
}

func (r *Resolver) Connect() error {
	if r.UserClient == nil {
		uconn, err := grpc.Dial(":"+r.UserPort, grpc.WithInsecure())
		if err != nil {
			return err
		}
		uclient := upb.NewUserServiceClient(uconn)
		r.UserClient = uclient
	}

	if r.ProjectClient == nil {
		pconn, err := grpc.Dial(":"+r.ProjectPort, grpc.WithInsecure())
		if err != nil {
			return err
		}
		pclient := ppb.NewProjectServiceClient(pconn)
		r.ProjectClient = pclient
	}

	return nil
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// AllProjects handle allProjects query
func (r *queryResolver) AllProjects(ctx context.Context) ([]project.Project, error) {
	if r.ProjectClient == nil {
		if err := r.Connect(); err != nil {
			log.Print(err)
			return nil, nil // ignore error as server is down
		}
	}
	stream, err := r.ProjectClient.AllProjects(ctx, &ppb.ProjectsRequest{})
	if err != nil {
		r.ProjectClient = nil
		log.Print(err)
		return nil, nil
	}
	projects := make([]project.Project, 0, 10)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		projects = append(projects, project.FromGRPC(in))
	}
	return projects, nil
}

// AllUsers handle allUsers query
func (r *queryResolver) AllUsers(ctx context.Context) ([]user.User, error) {
	if r.UserClient == nil {
		if err := r.Connect(); err != nil {
			log.Print(err)
			return nil, nil // ignore error as server is down
		}
	}
	stream, err := r.UserClient.AllUsers(ctx, &upb.UsersRequest{})
	if err != nil {
		r.UserClient = nil
		log.Print(err)
		return nil, nil
	}
	users := make([]user.User, 0, 10)
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		users = append(users, user.FromGRPC(in))
	}
	return users, nil
}

package graphql

import (
	"context"
	"exo/project"
	ppb "exo/project/pb"
	"exo/user"
	upb "exo/user/pb"
	"io"
)

type Resolver struct {
	ProjectClient ppb.ProjectServiceClient
	UserClient    upb.UserServiceClient
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

// AllProjects handle allProjects query
func (r *queryResolver) AllProjects(ctx context.Context) ([]project.Project, error) {
	stream, err := r.ProjectClient.AllProjects(ctx, &ppb.ProjectsRequest{})
	if err != nil {
		return nil, err
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
	stream, err := r.UserClient.AllUsers(ctx, &upb.UsersRequest{})
	if err != nil {
		return nil, err
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

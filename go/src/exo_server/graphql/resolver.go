package graphql

import (
	"context"
	"exo/project"
	"exo/user"
)

type Resolver struct{}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) AllProjects(ctx context.Context) ([]project.Project, error) {
	panic("needs gRPC")
}
func (r *queryResolver) AllUsers(ctx context.Context) ([]user.User, error) {
	panic("needs gRPC")
}

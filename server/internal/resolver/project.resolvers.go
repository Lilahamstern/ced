package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lilahamstern/ced/server/internal/graph/generated"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/project"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreatProjectInput) (*model.Project, error) {
	return project.Save(input)
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return project.GetAll()
}

func (r *queryResolver) Project(ctx context.Context, id int64) (*model.Project, error) {
	return project.Get(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

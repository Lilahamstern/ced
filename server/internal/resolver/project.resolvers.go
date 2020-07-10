package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/lilahamstern/ced/server/internal/graph/generated"
	"github.com/lilahamstern/ced/server/internal/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreateProjectInput) (*model.Project, error) {
	return r.services.Project.Save(input)
}

func (r *mutationResolver) DeleteProject(ctx context.Context, input *model.DeleteProjectInput) (bool, error) {
	return r.services.Project.Delete(input)
}

func (r *projectResolver) Versions(ctx context.Context, obj *model.Project) ([]*model.Version, error) {
	return r.services.Version.GetByProjectId(obj.ID)
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return r.services.Project.GetAll()
}

func (r *queryResolver) Project(ctx context.Context, id int64) (*model.Project, error) {
	return r.services.Project.Get(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{resolver} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{resolver} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{resolver} }

type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

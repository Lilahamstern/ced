package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/lilahamstern/ced/server/internal/graph/generated"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/pkg/service/project"
	"github.com/lilahamstern/ced/server/pkg/service/version"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreatProjectInput) (*model.Project, error) {
	return project.Save(input)
}

func (r *projectResolver) Versions(ctx context.Context, obj *model.Project) ([]*model.Version, error) {
	return version.GetByProjectId(obj.ID)
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return project.GetAll()
}

func (r *queryResolver) Project(ctx context.Context, id int64) (*model.Project, error) {
	return project.Get(id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Project returns generated.ProjectResolver implementation.
func (r *Resolver) Project() generated.ProjectResolver { return &projectResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type projectResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *projectResolver) Version(ctx context.Context, obj *model.Project) ([]*model.Version, error) {
	return version.GetByProjectId(obj.ID)
}

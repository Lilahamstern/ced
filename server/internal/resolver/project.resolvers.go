package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/lilahamstern/ced/server/internal/graph/model"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreatProjectInput) (*model.Project, error) {
	return r.ProjectService.Save(input)
}

func (r *projectResolver) Versions(ctx context.Context, obj *model.Project) ([]*model.Version, error) {
	return r.VersionService.GetByProjectId(obj.ID)
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	return r.ProjectService.GetAll()
}

func (r *queryResolver) Project(ctx context.Context, id int64) (*model.Project, error) {
	return r.ProjectService.Get(id)
}

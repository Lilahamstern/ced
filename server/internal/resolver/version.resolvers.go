package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/lilahamstern/ced/server/internal/graph/model"
)

func (r *mutationResolver) CreateVersion(ctx context.Context, input *model.CreateVersionInput) (*model.Version, error) {
	return r.VersionService.Save(input)
}

func (r *mutationResolver) DeleteVersion(ctx context.Context, input *model.DeleteVersionInput) (bool, error) {
	return r.VersionService.Delete(input)
}

func (r *queryResolver) Version(ctx context.Context, id string) (*model.Version, error) {
	return r.VersionService.GetById(id)
}

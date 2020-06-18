package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lilahamstern/ced/server/internal/graph/model"
)

func (r *mutationResolver) CreateVersion(ctx context.Context, input *model.CreateVersionInput) (*model.Version, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Version(ctx context.Context, id string) (*model.Version, error) {
	panic(fmt.Errorf("not implemented"))
}

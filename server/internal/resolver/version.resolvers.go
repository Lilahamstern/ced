package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lilahamstern/ced/server/internal/model/graph"
)

func (r *mutationResolver) CreateVersion(ctx context.Context, input *graph.CreateVersionInput) (*graph.Version, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Version(ctx context.Context, id string) (*graph.Version, error) {
	panic(fmt.Errorf("not implemented"))
}

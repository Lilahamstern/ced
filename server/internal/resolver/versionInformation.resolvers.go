package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lilahamstern/ced/server/internal/graph/model"
)

func (r *mutationResolver) CreateVersionInformation(ctx context.Context, input *model.CreateVersionInformationInput) (*model.VersionInformation, error) {
	return r.services.VersionInformation.Save(input)
}

func (r *mutationResolver) DeleteVersionInformation(ctx context.Context, input *model.DeleteVersionInformationInput) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

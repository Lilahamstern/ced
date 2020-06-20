package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lilahamstern/ced/server/internal/graph/generated"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/version"
)

func (r *mutationResolver) CreateVersion(ctx context.Context, input *model.CreateVersionInput) (*model.Version, error) {
	return version.Save(input)
}

func (r *queryResolver) Version(ctx context.Context, id string) (*model.Version, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *versionResolver) ProjectID(ctx context.Context, obj *model.Version) (*int64, error) {
	return &obj.ProjectId, nil
}

func (r *versionResolver) InformationID(ctx context.Context, obj *model.Version) (string, error) {
	return obj.InformationId, nil
}

// Version returns generated.VersionResolver implementation.
func (r *Resolver) Version() generated.VersionResolver { return &versionResolver{r} }

type versionResolver struct{ *Resolver }

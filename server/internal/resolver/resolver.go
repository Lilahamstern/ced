package resolver

import (
	"github.com/lilahamstern/ced/server/internal/graph/generated"
	"github.com/lilahamstern/ced/server/pkg/service"
)

//go:generate go run github.com/99designs/gqlgen

var resolver *Resolver

type Resolver struct {
	ProjectService service.ProjectService
	VersionService service.VersionService
}

func NewResolver(projectService service.ProjectService, versionService service.VersionService) {
	resolver = &Resolver{
		projectService,
		versionService,
	}
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

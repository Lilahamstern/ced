package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/lilahamstern/ced/server/internal/graph/generated"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/projects"
)

func (r *mutationResolver) CreateProject(ctx context.Context, input *model.CreatProjectInput) (*model.Project, error) {
	id, err := strconv.ParseInt(input.ID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}
	var project = projects.Project{
		ID: id,
	}

	project, err = project.Save()
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("internal server error")
	}

	return project.ToGraphModel(), nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var resProjects []*model.Project
	dbProjects := projects.GetAll()

	for _, project := range dbProjects {
		resProjects = append(resProjects, project.ToGraphModel())
	}

	return resProjects, nil
}

func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	project, err := projects.Get(id)
	if err != nil {
		return nil, err
	}

	return project.ToGraphModel(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

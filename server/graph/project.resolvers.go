package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/lilahamstern/ced/server/graph/generated"
	"github.com/lilahamstern/ced/server/graph/model"
	"github.com/lilahamstern/ced/server/internal/projects"
	"log"
	"strconv"
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

	return project.ToGraphProject(), nil
}

func (r *queryResolver) Projects(ctx context.Context) ([]*model.Project, error) {
	var respRojects []*model.Project
	dbProjects := projects.GetAll()

	for _, project := range dbProjects {
		respRojects = append(respRojects, project.ToGraphProject())
	}

	return respRojects, nil
}

func (r *queryResolver) Project(ctx context.Context, id string) (*model.Project, error) {
	project, err := projects.Get(id)
	if err != nil {
		return nil, err
	}

	return project.ToGraphProject(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package project

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/ced/server/internal/pkg/errors"
	"log"
)

func Save(input *model.CreatProjectInput) (*model.Project, error) {
	var project = database.Project{
		ID: input.ID,
	}

	if project.Exists() {
		return nil, &AlreadyExists{id: project.ID}
	}

	project, err := project.Save()
	if err != nil {
		log.Println(err)
		return nil, &errors.InternalServerError{}
	}

	return project.ToGraphModel(), nil
}

func GetAll() ([]*model.Project, error) {
	var projects []*model.Project
	dbProjects := database.GetAllProjects()

	for _, project := range dbProjects {
		projects = append(projects, project.ToGraphModel())
	}

	return projects, nil
}

func Get(id int64) (*model.Project, error) {
	project, err := database.GetProject(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFound{Id: id}
		}
		return nil, &errors.InternalServerError{}
	}

	return project.ToGraphModel(), nil
}

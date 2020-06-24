package project

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/pkg/errors"
	"github.com/lilahamstern/ced/server/internal/pkg/repository"
	"log"
)

func Save(input *model.CreatProjectInput) (*model.Project, error) {
	var project = repository.Project{
		ID: input.ID,
	}

	if project.Exists() {
		return nil, &AlreadyExists{id: project.ID}
	}

	err := project.Save()
	if err != nil {
		log.Println(err)
		return nil, &errors.InternalServerError{}
	}

	return project.ToGraphModel(), nil
}

func GetAll() ([]*model.Project, error) {
	var projects []repository.Project
	err := repository.GetAllProjects(&projects)
	if err != nil {
		return nil, err
	}
	var res []*model.Project
	for _, version := range projects {
		res = append(res, version.ToGraphModel())
	}

	return res, nil
}

func Get(id int64) (*model.Project, error) {
	project := repository.Project{
		ID: id,
	}
	err := project.GetProject()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFound{Id: id}
		}
		return nil, &errors.InternalServerError{}
	}

	return project.ToGraphModel(), nil
}

package project

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(input model.CreateProject) (domain.Project, error)
}

type Repo struct {
	DB *sql.DB
}

func (r Repo) Save(input model.CreateProject) (domain.Project, error) {
	const op errors.Op = "repo.project.save"
	query := "INSERT INTO projects(id) VALUES ($1) RETURNING createdat, updatedat"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return domain.Project{}, errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}
	defer stmt.Close()

	var project domain.Project
	err = stmt.QueryRow(input.ID).Scan(&project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		return domain.Project{}, errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}
	project.ID = input.ID

	return project, nil
}

func New(db *sql.DB) Repository {
	return &Repo{
		db,
	}
}

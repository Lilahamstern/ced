package project

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	log "github.com/sirupsen/logrus"
)

type Repository interface {
	Save(input model.CreateProject) error
	GetAll() ([]domain.Project, error)
}

type Repo struct {
	DB *sqlx.DB
}

func (r *Repo) GetAll() ([]domain.Project, error) {
	const op errors.Op = "repo.project.getAll"
	query := "SELECT * FROM get_projects_with_latest_version_view"

	stmt, err := r.DB.Preparex(query)
	if err != nil {
		return nil, r.Error(op, err)
	}

	rows, err := stmt.Queryx()
	if err != nil {
		return nil, r.Error(op, err)
	}

	var projects []domain.Project
	for rows.Next() {
		var project domain.Project
		err := rows.Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt,
			&project.Version.ID, &project.Version.OrderID, &project.Version.Title, &project.Version.Description,
			&project.Version.Manager, &project.Version.Client, &project.Version.Sector,
			&project.Version.CreatedAt, &project.Version.UpdatedAt)
		if err != nil {
			return nil, r.Error(op, err)
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// Save : saves record of project and version to database
func (r *Repo) Save(input model.CreateProject) error {
	const op errors.Op = "repo.project.save"
	tx, err := r.DB.Beginx()
	if err != nil {
		return r.Error(op, err)
	}

	projectStmt, err := tx.Preparex(`INSERT INTO project(id) VALUES ($1)`)
	if err != nil {
		return r.Error(op, err)
	}

	_, err = projectStmt.Exec(input.ID)
	if err != nil {
		return r.Error(op, err)
	}
	projectStmt.Close()

	versionStmt, err := tx.Preparex(`INSERT INTO version(projectid, orderid, title, description, manager, client, sector) 
				VALUES($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return r.Error(op, err)
	}

	_, err = versionStmt.Exec(input.ID, input.Version.OrderID, input.Version.Title, input.Version.Description,
		input.Version.Manager, input.Version.Client, input.Version.Sector)
	if err != nil {
		return r.Error(op, err)
	}
	versionStmt.Close()

	err = tx.Commit()
	if err != nil {
		return r.Error(op, err)
	}

	return nil
}

func New(db *sqlx.DB) Repository {
	return &Repo{
		db,
	}
}

func (r Repo) Error(op errors.Op, err error) error {
	return errors.E(op, err, errors.KindInternalServer, log.WarnLevel)
}

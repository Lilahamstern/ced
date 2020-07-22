package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"
	log "github.com/sirupsen/logrus"
)

type ProjectRepository interface {
	Save(input request.CreateProject) error
	GetAll() ([]domain.Project, error)
	Get(id string) (*domain.Project, error)
}

type project struct {
	db *sqlx.DB
}

func (r *project) Get(id string) (*domain.Project, error) {
	const op errors.Op = "repo.project.get"
	query := "SELECT id AS project_id, createdat AS project_createdat, updatedat AS project_updatedAt FROM project WHERE id = $1"

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, r.Error(op, err)
	}

	rows, err := stmt.Queryx(id)
	if err != nil {
		return nil, r.Error(op, err)
	}

	var project *domain.Project = nil
	for rows.Next() {
		err := rows.StructScan(&project)
		if err != nil {
			rows.Close()
			return nil, r.Error(op, err)
		}
	}

	return project, nil
}

func (r *project) GetAll() ([]domain.Project, error) {
	const op errors.Op = "repo.project.getAll"
	query := "SELECT * FROM get_projects_with_latest_version_view"

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, r.Error(op, err)
	}
	defer stmt.Close()

	rows, err := stmt.Queryx()
	if err != nil {
		return nil, r.Error(op, err)
	}

	var projects []domain.Project
	for rows.Next() {
		var project domain.Project
		err := rows.StructScan(&project)

		if err != nil {
			rows.Close()
			return nil, r.Error(op, err)
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// Save : saves record of project and version to database
func (r *project) Save(input request.CreateProject) error {
	const op errors.Op = "repo.project.save"

	stmt, err := r.db.Preparex("CALL createProjectWithVersion($1, $2, $3, $4, $5, $6, $7);")
	if err != nil {
		return r.Error(op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(input.ID, input.Version.OrderID, input.Version.Title,
		input.Version.Description, input.Version.Manager, input.Version.Client, input.Version.Sector)
	if err != nil {
		return r.Error(op, err)
	}
	return nil
}

func (r project) Error(op errors.Op, err error) error {
	return errors.E(op, err, errors.KindError, log.WarnLevel)
}

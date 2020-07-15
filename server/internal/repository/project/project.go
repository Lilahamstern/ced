package project

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(input model.CreateProject) error
}

type Repo struct {
	DB *sqlx.DB
}

// Save :
func (r *Repo) Save(input model.CreateProject) error {
	const op errors.Op = "repo.project.save"
	tx, err := r.DB.Beginx()
	if err != nil {
		return errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}

	projectStmt, err := tx.Preparex(`INSERT INTO project(id) VALUES ($1)`)
	if err != nil {
		return errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}

	_, err = projectStmt.Exec(input.ID)
	if err != nil {
		return errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}
	projectStmt.Close()

	versionStmt, err := tx.Preparex(`INSERT INTO version(projectid, orderid, title, description, manager, client, sector) 
				VALUES($1, $2, $3, $4, $5, $6, $7)`)
	if err != nil {
		return errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}

	_, err = versionStmt.Exec(input.ID, input.Version.OrderID, input.Version.Title, input.Version.Description,
		input.Version.Manager, input.Version.Client, input.Version.Sector)
	if err != nil {
		return errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}
	versionStmt.Close()

	err = tx.Commit()
	if err != nil {
		return errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}

	return nil
}

func New(db *sqlx.DB) Repository {
	return &Repo{
		db,
	}
}

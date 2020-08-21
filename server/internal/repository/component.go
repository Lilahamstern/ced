package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"
	log "github.com/sirupsen/logrus"
)

type ComponentRepository interface {
	Save(input []*request.CreateComponent) error
}

type component struct {
	db *sqlx.DB
}

// Save : saves record of project and version to database
func (r *component) Save(input []*request.CreateComponent) error {
	const op errors.Op = "repo.project.save"

	query := "INSERT INTO components(versionId, compid, name, profile, material, co, level, type) VALUES ($1, $2, $3, $4, $5, $6, $7);"

	res, err := r.db.NamedExec(query, input)
	fmt.Println(res)
	if err != nil {
		return r.Error(op, err)
	}
	return nil
}

func (r component) Error(op errors.Op, err error) error {
	return errors.E(op, err, errors.KindError, log.WarnLevel)
}

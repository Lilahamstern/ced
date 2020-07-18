package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"

	log "github.com/sirupsen/logrus"
)

type VersionRepository interface {
	GetVersionsByProjectId(id string) ([]domain.Version, error)
}

type version struct {
	db *sqlx.DB
}

func (r *version) GetVersionsByProjectId(id string) ([]domain.Version, error) {
	const op errors.Op = "repo.version.getVersionsByProjectId"

	query := `SELECT id AS version_id, orderid AS version_order_id, title AS version_title, 
       description AS version_description, manager AS version_manager, client AS version_client, 
       sector AS version_sector, createdat AS version_createdat, updatedat AS version_updatedat 
	FROM version WHERE projectid = $1`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return nil, r.Error(op, err)
	}
	defer stmt.Close()

	versions := []domain.Version{}
	err = stmt.Select(&versions, id)
	if err != nil {
		return nil, r.Error(op, err)
	}

	return versions, nil
}

func (r version) Error(op errors.Op, err error) error {
	return errors.E(op, err, errors.KindError, log.WarnLevel)
}

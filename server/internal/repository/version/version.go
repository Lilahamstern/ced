package version

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"

	log "github.com/sirupsen/logrus"
)

type Repository interface {
	GetVersionsByProjectId(id string) ([]domain.Version, error)
}

type Repo struct {
	DB *sqlx.DB
}

func (r *Repo) GetVersionsByProjectId(id string) ([]domain.Version, error) {
	const op errors.Op = "repo.version.getVersionsByProjectId"
	query := `SELECT id AS version_id, orderid AS version_order_id, title AS version_title, 
       description AS version_description, manager AS version_manager, client AS version_client, 
       sector AS version_sector, createdat AS version_createdat, updatedat AS version_updatedat 
	FROM version WHERE projectid = $1`
	stmt, err := r.DB.Preparex(query)
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

func New(db *sqlx.DB) Repository {
	return &Repo{
		db,
	}
}

func (r Repo) Error(op errors.Op, err error) error {
	return errors.E(op, err, errors.KindInternalServer, log.WarnLevel)
}

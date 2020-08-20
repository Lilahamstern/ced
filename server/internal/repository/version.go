package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model/request"

	log "github.com/sirupsen/logrus"
)

type VersionRepository interface {
	GetVersionsByProjectID(id string) ([]domain.Version, error)
	Save(projectID string, input request.CreateVersion) error
}

type version struct {
	db *sqlx.DB
}

func (r *version) GetVersionsByProjectID(id string) ([]domain.Version, error) {
	const op errors.Op = "repo.version.getVersionsByProjectId"

	query := `SELECT id AS version_id,
				   orderid AS version_order_id,
				   title AS version_title,
				   description AS version_description,
				   manager AS version_manager,
				   client AS version_client,
				   sector AS version_sector,
				   createdat AS version_createdat,
				   updatedat AS version_updatedat,
				   coalesce(c1.total, 0) AS version_co
			 FROM VERSION AS v
			 LEFT JOIN
			   (SELECT coalesce(SUM(c.co), 0) AS total,
					   c.versionid
			    FROM components AS c
			    GROUP BY c.versionid) AS c1 ON c1.versionid = v.id
			 WHERE projectid = $1 ORDER BY updatedat DESC;`

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

// Save : saves record of version and  to database
func (r *version) Save(projectID string, input request.CreateVersion) error {
	const op errors.Op = "repo.version.save"

	query := `INSERT INTO version(projectid, orderid, title, description, manager, client, sector)
    VALUES ($1, $2, $3, $4, $5, $6, $7);`

	stmt, err := r.db.Preparex(query)
	if err != nil {
		return r.Error(op, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(projectID, input.OrderID, input.Title,
		input.Description, input.Manager, input.Client, input.Sector)
	if err != nil {
		return r.Error(op, err)
	}
	return nil
}

func (r version) Error(op errors.Op, err error) error {
	return errors.E(op, err, errors.KindError, log.WarnLevel)
}

package version

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/pkg/errors"
	"github.com/lilahamstern/ced/server/pkg/model"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Save(projectId int64, input model.CreateVersion) (*domain.Version, error)
}

type Repo struct {
	DB *sql.DB
}

func (r Repo) Save(projectId int64, input model.CreateVersion) (*domain.Version, error) {
	const op errors.Op = "repo.version.save"
	query := "INSERT INTO version(projectid, orderid, title, description, manager, client, sector) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id, createdat, updatedat"
	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return &domain.Version{}, errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}
	defer stmt.Close()

	version := domain.NewVersion(input)

	err = stmt.QueryRow(projectId, version.OrderID, version.Title, version.Description, version.Manager, version.Client, version.Sector).Scan(&version.ID, &version.CreatedAt, &version.UpdatedAt)
	if err != nil {
		return &domain.Version{}, errors.E(op, err, errors.KindInternalServer, logrus.WarnLevel)
	}

	return version, nil
}

func New(db *sql.DB) Repository {
	return &Repo{
		db,
	}
}

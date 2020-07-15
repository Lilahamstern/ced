package information

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/pkg/model"
)

type Repository interface {
	Save(input model.CreateInformation) (domain.Version, error)
}

type Repo struct {
	DB *sql.DB
}

func (r Repo) Save(input model.CreateInformation) (domain.Version, error) {
	panic("implement me")
}

func New(db *sql.DB) Repository {
	return &Repo{
		db,
	}
}

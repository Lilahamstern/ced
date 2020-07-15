package version

import (
	"github.com/jmoiron/sqlx"
)

type Repository interface {
}

type Repo struct {
	DB *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &Repo{
		db,
	}
}

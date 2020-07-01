package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/pkg/domain"
	"log"
)

type VersionInformationRepository interface {
	Save(info *domain.VersionInformation) error
}

type versionInfomrationRepository struct {
	db *sql.DB
}

func (repo versionInfomrationRepository) Save(info *domain.VersionInformation) error {
	query := "INSERT INTO versioninformation(orderid, title, description, manager, client, sector) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, createdat, updatedat"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(info.OrderID, info.Title, info.Description, info.Manager, info.Client, info.Sector).Scan(
		&info.ID,
		&info.CreatedAt,
		&info.UpdatedAt,
	)

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func NewVersionInformationRepository(db *sql.DB) VersionInformationRepository {
	return &versionInfomrationRepository{
		db,
	}
}

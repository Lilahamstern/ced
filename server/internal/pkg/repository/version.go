package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/pkg/domain"
	"log"
)

type VersionRepository interface {
	Save(version *domain.Version) error
	GetVersionByProjectId(id int64, versions *[]domain.Version) error
	GetVersionById(version domain.Version) error
}

type versionRepository struct {
	db *sql.DB
}

func (repo versionRepository) Save(version *domain.Version) error {
	query := "INSERT INTO versions(projectid, informationid) VALUES ($1, $2) RETURNING id, projectid, informationid, createdat, updatedat"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(version.ProjectId, version.InformationId).Scan(
		&version.ID,
		&version.ProjectId,
		&version.InformationId,
		&version.CreatedAt,
		&version.UpdatedAt)

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (repo versionRepository) GetVersionByProjectId(id int64, versions *[]domain.Version) error {
	query := "SELECT id, projectid, informationid, createdat, updatedat FROM versions WHERE projectid = $1"
	stmt, err := repo.db.Prepare(query)
	log.Println(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var version domain.Version
		err := rows.Scan(&version.ID, &version.ProjectId, &version.InformationId, &version.CreatedAt, &version.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		*versions = append(*versions, version)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

func (repo versionRepository) GetVersionById(version domain.Version) error {
	query := "SELECT projectid, informationid, createdat, updatedat FROM versions WHERE id = $1"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(version.ID).Scan(
		&version.ProjectId,
		&version.InformationId,
		&version.CreatedAt,
		&version.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Fatal(err)
	}

	return nil
}

func NewVersionRepository(db *sql.DB) VersionRepository {
	return &versionRepository{
		db: db,
	}
}

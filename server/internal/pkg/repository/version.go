package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/pkg/model"
	"log"
)

type VersionRepository struct {
	db *sql.DB
}

func (repo VersionRepository) Save(version *model.Version) error {
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

func (repo VersionRepository) GetVersionByProjectId(id int64, versions *[]model.Version) error {
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
		var version model.Version
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

func (repo VersionRepository) GetVersionById(version model.Version) error {
	query := "SELECT projectid, informationid, createdat, updatedat FROM versions WHERE id = $1"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(version.ID).Scan(
		&version.ID,
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

func NewVersionRepository(db *sql.DB) *VersionRepository {
	return &VersionRepository{
		db: db,
	}
}

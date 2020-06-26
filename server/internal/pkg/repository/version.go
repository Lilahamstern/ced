package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"log"
	"time"
)

// Version model for database
type Version struct {
	ID            uuid.UUID `json:"id"`
	ProjectId     int64     `json:"project_id"`
	InformationId uuid.UUID `json:"information_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (v Version) ToGraphModel() *model.Version {
	return &model.Version{
		ID:            v.ID.String(),
		ProjectId:     v.ProjectId,
		InformationId: v.InformationId.String(),
		CreatedAt:     v.CreatedAt.String(),
		UpdatedAt:     v.UpdatedAt.String(),
	}
}

func (v *Version) Save() error {
	query := "INSERT INTO versions(projectid, informationid) VALUES ($1, $2) RETURNING id, projectid, informationid, createdat, updatedat"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(v.ProjectId, v.InformationId).Scan(
		&v.ID,
		&v.ProjectId,
		&v.InformationId,
		&v.CreatedAt,
		&v.UpdatedAt)

	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func GetVersionByProjectId(id int64, versions *[]Version) error {
	query := "SELECT id, projectid, informationid, createdat, updatedat FROM versions WHERE projectid = $1"
	stmt, err := database.DB.Prepare(query)
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
		var version Version
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

func (v *Version) GetVersionById() error {
	query := "SELECT id, projectid, informationid, createdat, updatedat FROM versions WHERE id = $1"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(v.ID).Scan(&v.ID, &v.ProjectId, &v.InformationId, &v.CreatedAt, &v.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Fatal(err)
	}

	return nil
}

package database

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
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

func (v Version) Save() (Version, error) {
	query := "INSERT INTO versions(projectid, informationid) VALUES ($1, $2) RETURNING id, projectid, informationid, createdat, updatedat"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(v.ProjectId, v.InformationId).Scan(&v.ID, &v.ProjectId, &v.InformationId, &v.CreatedAt, &v.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return v, nil
}

func GetVersionById(id string) (Version, error) {
	query := "SELECT id, projectid, informationid, createdat, updatedat FROM versions WHERE id = $1"
	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var version Version
	err = stmt.QueryRow(id).Scan(&version.ID, &version.ProjectId, &version.InformationId, &version.CreatedAt, &version.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Version{}, err
		}
		log.Fatal(err)
	}

	return version, nil
}

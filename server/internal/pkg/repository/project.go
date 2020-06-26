package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"log"
	"time"
)

// Project struct of database model
type Project struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToGraphModel will map Project to Graphql model of project and return it as pointer
func (p Project) ToGraphModel() *model.Project {
	return &model.Project{
		ID:        p.ID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

// Save will save project to db, if project.ID isn't unique error will be returned
func (p *Project) Save() error {
	query := "INSERT INTO projects(id) VALUES ($1) RETURNING id, createdat, updatedat"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(p.ID).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (p Project) Exists() bool {
	query := "SELECT EXISTS(SELECT 1 FROM projects p WHERE p.id=$1)"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(p.ID).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func ProjectExistsById(id int64) bool {
	project := Project{
		ID: id,
	}

	return project.Exists()
}

// GetAll will fetch all projects from database
func GetAllProjects(projects *[]Project) error {
	query := "SELECT p.id, p.createdat, p.updatedat FROM projects p"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		*projects = append(*projects, project)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// Get will query one project from database with provided id
func (p *Project) GetProject() error {
	query := "SELECT p.id, p.createdat, p.updatedat FROM projects p WHERE id = $1"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(p.ID).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Fatal(err)
	}

	return nil
}

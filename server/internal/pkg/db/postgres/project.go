package database

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/graph/model"
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
func (p Project) Save() (Project, error) {
	query := "INSERT INTO projects(id) VALUES ($1) RETURNING id, createdat, updatedat"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(p.ID).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return p, nil
}

func (p Project) Exists() bool {
	query := "SELECT EXISTS(SELECT 1 FROM projects p WHERE p.id=$1)"
	stmt, err := DB.Prepare(query)
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
func GetAllProjects() []Project {
	query := "SELECT p.id, p.createdat, p.updatedat FROM projects p"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	var projects []Project
	for rows.Next() {
		var project Project
		err := rows.Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return projects
}

// Get will query one project from database with provided id
func GetProject(id int64) (Project, error) {
	query := "SELECT p.id, p.createdat, p.updatedat FROM projects p WHERE id = $1"

	stmt, err := DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var project Project
	err = stmt.QueryRow(id).Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Project{}, err
		}
		log.Fatal(err)
	}

	return project, nil
}

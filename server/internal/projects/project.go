package projects

import (
	"database/sql"
	"fmt"
	"github.com/lilahamstern/ced/server/graph/model"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"log"
	"strconv"
	"time"
)

// Project struct used to CRUD projects
type Project struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToGrpahProject will map Project to Graphql model of project and return it as pointer
func (project Project) ToGraphProject() *model.Project {
	return &model.Project{
		ID:        strconv.FormatInt(project.ID, 10),
		CreatedAt: project.CreatedAt.String(),
		UpdatedAt: project.UpdatedAt.String(),
	}
}

// Save will save project to db, if project.ID isn't unique error will be returned
func (project Project) Save() (Project, error) {
	query := "INSERT INTO projects(id) VALUES ($1) RETURNING id, createdat, updatedat"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(project.ID).Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return project, nil
}

// GetAll will fetch all projects from database
func GetAll() []Project {
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
func Get(id string) (Project, error) {
	query := "SELECT p.id, p.createdat, p.updatedat FROM projects p WHERE id = $1"

	stmt, err := database.DB.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var project Project
	err = stmt.QueryRow(id).Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Project{}, fmt.Errorf("project not found id: %v", id)
		}
		log.Fatal(err)
	}

	return project, nil
}

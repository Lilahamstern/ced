package repository

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/pkg/model"
	"log"
)

// ProjectRepository struct of database model
type ProjectRepository struct {
	db *sql.DB
}

// Save will save project to db, if project.ID isn't unique error will be returned
func (repo ProjectRepository) Save(project *model.Project) error {
	query := "INSERT INTO projects(id) VALUES ($1) RETURNING id, createdat, updatedat"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(project.ID).Scan(&project.ID, &project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

func (repo ProjectRepository) ExistsById(id int64) bool {
	query := "SELECT EXISTS(SELECT 1 FROM projects p WHERE p.id=$1)"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	var exists bool
	err = stmt.QueryRow(id).Scan(&exists)
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

// GetAll will fetch all projects from database
func (repo ProjectRepository) GetAllProjects(projects *[]model.Project) error {
	query := "SELECT p.id, p.createdat, p.updatedat FROM projects p"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var project model.Project
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
func (repo ProjectRepository) GetProject(project *model.Project) error {
	query := "SELECT p.createdat, p.updatedat FROM projects p WHERE id = $1"

	stmt, err := repo.db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	err = stmt.QueryRow(project.ID).Scan(&project.CreatedAt, &project.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		log.Fatal(err)
	}

	return nil
}

// NewProjectRepository
func NewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

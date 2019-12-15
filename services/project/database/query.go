package database

import (
	"fmt"
	"project/models"
)

// CreateProject stores project data in db
func (dc *DBClient) CreateProject(project models.Project) (models.Project, error) {
	e := dc.db.Create(&project).Error

	if e != nil {
		return models.Project{}, e
	}

	return project, nil
}

// QueryAllProjects will return a list of all projects
func (dc *DBClient) QueryAllProjects(limit interface{}) ([]models.Project, []error) {
	var projects []models.Project
	errs := dc.db.Limit(limit).Find(&projects).GetErrors()
	fmt.Println("Skit")
	if len(errs) >= 1 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return projects, errs
	}

	return projects, nil
}

// SearchProjects searching projects from database
func (dc *DBClient) SearchProjects(search string, limit interface{}) ([]models.Project, []error) {
	var projects []models.Project
	errs := dc.db.Limit(limit).Where("lower(id) like ? OR lower(name) like ?", "%"+search+"%", "%"+search+"%").Find(&projects).GetErrors()

	if len(errs) > 1 {
		for _, err := range errs {
			fmt.Println(err)
		}
		return projects, errs
	}

	return projects, nil
}

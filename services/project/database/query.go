package database

import "project/models"

// CreateProject stores project data in db
func (dc *DBClient) CreateProject(project models.Project) (models.Project, error) {
	e := dc.db.Create(&project).Error

	if e != nil {
		return models.Project{}, e
	}

	return project, nil
}

// QueryAllProjects will return a list of all projects
func (dc *DBClient) QueryAllProjects(limit interface{}) ([]models.Project, error) {
	var projects []models.Project
	err := dc.db.Limit(limit).Find(&projects).Error
	if err != nil {
		return projects, err
	}

	return projects, nil
}

// SearchProjects searching projects from database
func (dc *DBClient) SearchProjects(search string, limit interface{}) ([]models.Project, error) {
	var projects []models.Project
	err := dc.db.Limit(limit).Where("lower(id) like ? OR lower(name) like ?", "%"+search+"%", "%"+search+"%").Find(&projects).Error

	if err != nil {
		return projects, err
	}

	return projects, nil
}

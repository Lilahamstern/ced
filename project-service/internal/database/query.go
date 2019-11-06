package database

import (
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
)

func (c *Client) QueryProject(projectId string) (model.Project, error) {
	project := model.Project{}

	err := c.db.Where("id = ?", projectId).Find(&project).Error

	if err != nil {
		return model.Project{}, err
	}

	return project, nil
}

func (c *Client) CreateProject(project model.Project) (model.Project, error) {
	e := c.db.Create(&project).Error

	if e != nil {
		return model.Project{}, e
	}

	return project, nil
}

func (c *Client) DeleteProject(projectId string) error {
	e := c.db.Delete(&model.Project{}, "id = ?", projectId).Error
	if e != nil {
		return e
	}

	return nil
}

func (c *Client) QueryAllProjects() ([]model.Project, error) {
	var projects []model.Project
	e := c.db.Find(&projects).Error

	if e != nil {
		return []model.Project{}, e
	}

	return projects, nil
}

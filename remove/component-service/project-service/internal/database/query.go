package database

import (
	"fmt"
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

func (c *Client) QueryAllProjects() (model.Projects, error) {
	var projects model.Projects
	e := c.db.Find(&projects).Limit(30).Error

	if e != nil {
		return projects, e
	}

	return projects, nil
}

func (c *Client) ProjectExists(project model.Project) error {
	var exists model.Project
	var e error

	e = c.db.Where("id = ?", project.Id).Find(&exists).Error
	if e == nil {
		return fmt.Errorf("id: %v, already exists", project.Id)
	}
	fmt.Println(e)

	e = c.db.Where("name = ?", project.Name).Find(&exists).Error
	if e == nil {
		return fmt.Errorf("name: %v, already exists", project.Name)
	}

	e = c.db.Where("model = ?", project.Model).Find(&exists).Error
	if e == nil {
		return fmt.Errorf("model: %v, already exists", project.Model)
	}

	return nil
}

func (c *Client) SearchProjects(search string) (model.Projects, error) {
	var projects model.Projects
	e := c.db.Where("lower(id) like ? OR lower(client) like ?", "%"+search+"%", "%"+search+"%").Find(&projects).Error

	if e != nil {
		return projects, e
	}

	return projects, nil
}

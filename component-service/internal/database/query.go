package database

import (
	"github.com/lilahamstern/bec-microservices/component-service/internal/model"
)

func (c *Client) QueryComponents(projectId string) (model.Components, error) {
	var components model.Components

	err := c.db.Where("project_id = ?", projectId).Find(&components).Error

	if err != nil {
		return components, err
	}

	return components, nil
}

func (c *Client) CreateComponent(component model.Component) error {
	e := c.db.Create(&component).Error

	if e != nil {
		return e
	}

	return nil
}

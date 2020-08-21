package domain

import (
	"time"

	"github.com/lilahamstern/ced/server/pkg/model/response"
)

// Component : Database model
type Component struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Profile   string    `json:"profile"`
	Material  string    `json:"material"`
	Co        float64   `json:"co"`
	Level     int64     `json:"level"`
	Type      string    `json:"type"`
	CreatedAt time.Time `db:"project_createdat"`
	UpdatedAt time.Time `db:"project_updatedat"`
}

// ToModel : will map database model of project to request model of project
func (c *Component) ToModel() *response.Component {
	return &response.Component{
		ID: c.ID,
		
	}
}

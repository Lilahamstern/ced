package domain

import (
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Project : Database model
type Project struct {
	ID        int64     `json:"id"`
	Version   Version   `json:"version"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToModel : will map database model of project to request model of project
func (p *Project) ToModel() *model.Project {
	return &model.Project{
		ID:        p.ID,
		Version:   p.Version.ToModel(),
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

// ToModel : maps slice of domain project to slice of response project model
func ToModel(projects []Project) []model.Project {
	var p []model.Project
	for _, v := range projects {
		p = append(p, *v.ToModel())
	}

	return p
}

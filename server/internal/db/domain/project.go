package domain

import (
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Project : Database model
type Project struct {
	ID int64 `db:"project_id"`
	Version
	CreatedAt time.Time `db:"project_createdat"`
	UpdatedAt time.Time `db:"project_updatedat"`
}

// ToModel : will map database model of project to request model of project
func (p *Project) ToModel() *model.Project {
	return &model.Project{
		ID:        p.ID,
		Version:   *p.Version.ToModel(),
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

// ToModel : maps slice of domain project to slice of response project model
func ToModel(projects []Project) []model.Project {
	if projects == nil {
		return []model.Project{}
	}

	var p []model.Project
	for _, v := range projects {
		p = append(p, *v.ToModel())
	}

	return p
}

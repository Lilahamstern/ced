package model

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"time"
)

// Project model
type Project struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToGraphModel will map ProjectRepository to Graphql model of project and return it as pointer
func (p Project) ToGraphModel() *model.Project {
	return &model.Project{
		ID:        p.ID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

// Version model for database
type Version struct {
	ID            uuid.UUID `json:"id"`
	ProjectId     int64     `json:"project_id"`
	InformationId uuid.UUID `json:"information_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (v Version) ToGraphModel() *model.Version {
	return &model.Version{
		ID:            v.ID.String(),
		ProjectId:     v.ProjectId,
		InformationId: v.InformationId.String(),
		CreatedAt:     v.CreatedAt.String(),
		UpdatedAt:     v.UpdatedAt.String(),
	}
}

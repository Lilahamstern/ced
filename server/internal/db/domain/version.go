package domain

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Version : Database model
type Version struct {
	ID            uuid.UUID `json:"id"`
	ProjectId     int64     `json:"project_id"`
	InformationId uuid.UUID `json:"information_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// ToModel : will map database model of version to request model of version
func (v *Version) ToModel() *model.Version {
	return &model.Version{
		ID:        v.ID,
		CreatedAt: v.CreatedAt.String(),
		UpdatedAt: v.UpdatedAt.String(),
	}
}

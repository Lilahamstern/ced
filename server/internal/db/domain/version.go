package domain

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Version : Database model
type Version struct {
	ID          *uuid.UUID `db:"version_id"`
	OrderID     *int64     `json:"orderId"`
	Title       *string    `json:"title"`
	Description *string    `json:"description"`
	Manager     *string    `json:"manager"`
	Client      *string    `json:"client"`
	Sector      *string    `json:"sector"`
	CreatedAt   *time.Time `db:"version_createdat"`
	UpdatedAt   *time.Time `db:"version_updatedat"`
}

// ToModel : maps database model of Version to request model of Version
func (i *Version) ToModel() *model.Version {
	if i.ID == nil {
		return &model.Version{
			ID: "",
		}
	}
	return &model.Version{
		ID:          i.ID.String(),
		OrderID:     *i.OrderID,
		Title:       *i.Title,
		Description: *i.Description,
		Manager:     *i.Manager,
		Client:      *i.Client,
		Sector:      *i.Sector,
		CreatedAt:   i.CreatedAt.String(),
		UpdatedAt:   i.UpdatedAt.String(),
	}
}

// NewVersion : Converts CreateVersion to Version model of domain
func NewVersion(s model.CreateVersion) *Version {
	return &Version{
		OrderID:     &s.OrderID,
		Title:       &s.Title,
		Description: &s.Description,
		Manager:     &s.Manager,
		Client:      &s.Client,
		Sector:      &s.Sector,
	}
}

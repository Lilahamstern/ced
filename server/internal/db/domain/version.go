package domain

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Version : Database model
type Version struct {
	ID          uuid.UUID `json:"-"`
	OrderID     int64     `json:"orderId"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Manager     string    `json:"manager"`
	Client      string    `json:"client"`
	Sector      string    `json:"sector"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ToModel : will map database model of Version to request model of Version
func (i *Version) ToModel() *model.Version {
	return &model.Version{
		ID:          i.ID,
		OrderID:     i.OrderID,
		Title:       i.Title,
		Description: *i.Description,
		Manager:     i.Manager,
		Client:      i.Client,
		Sector:      i.Sector,
		CreatedAt:   i.CreatedAt.String(),
		UpdatedAt:   i.UpdatedAt.String(),
	}
}

// NewVersion : Converts CreateVersion to Version model of domain
func NewVersion(s model.CreateVersion) *Version {
	return &Version{
		OrderID:     s.OrderID,
		Title:       s.Title,
		Description: &s.Description,
		Manager:     s.Manager,
		Client:      s.Client,
		Sector:      s.Sector,
	}
}

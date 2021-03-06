package domain

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/pkg/model/response"
	"time"
)

// Version : Database model
type Version struct {
	ID          *uuid.UUID `db:"version_id"`
	OrderID     *int64     `db:"version_order_id"`
	Title       *string    `db:"version_title"`
	Description *string    `db:"version_description"`
	Manager     *string    `db:"version_manager"`
	Client      *string    `db:"version_client"`
	Sector      *string    `db:"version_sector"`
	Co          int        `db:"version_co"`
	CreatedAt   *time.Time `db:"version_createdat"`
	UpdatedAt   *time.Time `db:"version_updatedat"`
}

// ToModel : maps database model of Version to request model of Version
func (i *Version) ToModel() *response.Version {
	if i.ID == nil {
		return &response.Version{
			ID: "",
		}
	}

	return &response.Version{
		ID:          i.ID.String(),
		OrderID:     *i.OrderID,
		Title:       *i.Title,
		Description: i.Description,
		Manager:     *i.Manager,
		Client:      *i.Client,
		Sector:      *i.Sector,
		Co:          i.Co,
		CreatedAt:   i.CreatedAt.String(),
		UpdatedAt:   i.UpdatedAt.String(),
	}
}

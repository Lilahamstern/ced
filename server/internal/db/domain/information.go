package domain

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Information : Database model
type Information struct {
	ID          uuid.UUID `json:"id"`
	OrderID     int64     `json:"orderId"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Manager     string    `json:"manager"`
	Client      string    `json:"client"`
	Sector      string    `json:"sector"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

// ToModel : will map database model of information to request model of information
func (i *Information) ToModel() *model.Information {
	return &model.Information{
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

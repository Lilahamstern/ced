package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Base model struct
type Base struct {
	UUID      uuid.UUID `json:"uuid" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate of data
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", uuid.New())
	return nil
}

package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// Base model struct
type Base struct {
	UID       uuid.UUID `json:"uid" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate of data
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UID", uuid.New())
	return nil
}

package models

import (
	"time"
)

// Base model struct
type Base struct {
	CreatedAt time.Time `ison:"createdAt"`
	UpdatedAt time.Time `ison:"updatedAt"`
}

// BeforeCreate of data
// func (base *Base) BeforeCreate(scope *gorm.Scope) error {
// 	scope.SetColumn("UID", uuid.New())
// 	return nil
// }

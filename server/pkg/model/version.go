package model

import (
	"github.com/google/uuid"
)

type (
	Version struct {
		ID        uuid.UUID
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

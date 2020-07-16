package model

import (
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type (
	Version struct {
		ID          string `json:"id,omitempty"`
		OrderID     int64  `json:"order_id,omitempty"`
		Title       string `json:"title,omitempty"`
		Description string `json:"desc,omitempty"`
		Manager     string `json:"manager,omitempty"`
		Client      string `json:"client,omitempty"`
		Sector      string `json:"sector,omitempty"`
		CreatedAt   string `json:"created_at,omitempty"`
		UpdatedAt   string `json:"updated_at,omitempty"`
	}

	CreateVersion struct {
		OrderID     int64  `json:"order_id"`
		Title       string `json:"title"`
		Description string `json:"desc"`
		Manager     string `json:"manager"`
		Client      string `json:"client"`
		Sector      string `json:"sector"`
	}
)

// Validate : Validates CreateInformation struct to match requirements to create record
func (s *CreateVersion) Validate() url.Values {
	rules := govalidator.MapData{
		"order_id": []string{"required"},
		"title":    []string{"required"},
		"manager":  []string{"required"},
		"client":   []string{"required"},
		"sector":   []string{"required"},
	}

	messages := govalidator.MapData{}

	opts := govalidator.Options{
		Data:     s,
		Rules:    rules,
		Messages: messages,
	}

	e := validation.Validate(opts)

	return e
}

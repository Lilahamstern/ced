package model

import (
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type (
	Information struct {
		ID          uuid.UUID `json:"id"`
		OrderID     int64     `json:"order_id"`
		Title       string    `json:"title"`
		Description string    `json:"desc"`
		Manager     string    `json:"manager"`
		Client      string    `json:"client"`
		Sector      string    `json:"sector"`
		CreatedAt   string    `json:"created_at"`
		UpdatedAt   string    `json:"updated_at"`
	}

	CreateInformation struct {
		OrderID     int64  `json:"order_id"`
		Title       string `json:"title"`
		Description string `json:"desc"`
		Manager     string `json:"manager"`
		Client      string `json:"client"`
		Sector      string `json:"sector"`
	}
)

func (s *CreateInformation) Validate() url.Values {
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

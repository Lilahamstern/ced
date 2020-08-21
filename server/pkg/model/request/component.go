package request

import (
	"net/url"

	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/thedevsaddam/govalidator"
)

type CreateComponent struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Profile  string  `json:"profile"`
	Material string  `json:"material"`
	Co       float64 `json:"co"`
	Level    int64   `json:"level"`
	Type     string  `json:"type"`
}

// Validate : Validates CreateInformation struct to match requirements to create record
func (s *CreateComponent) Validate() url.Values {
	rules := govalidator.MapData{
		"id":       []string{"required"},
		"name":     []string{"required"},
		"profile":  []string{"required"},
		"material": []string{"required"},
		"co":       []string{"required"},
		"level":    []string{"required"},
		"type":     []string{"required"},
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

package model

import (
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type (
	Project struct {
		ID        int64   `json:"id"`
		Version   Version `json:"version"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
	}

	CreateProject struct {
		ID      int64         `json:"id"`
		Version CreateVersion `json:"version"`
	}
)

// Validate : Validates CreateInformation struct to match requirements to create record
func (s *CreateProject) Validate() url.Values {
	rules := govalidator.MapData{
		"id": []string{"required", "min:9999999", "unique:project"},
	}

	messages := govalidator.MapData{
		"id": []string{"min:The id field cannot be less then 8 numbers"},
	}

	opts := govalidator.Options{
		Data:     s,
		Rules:    rules,
		Messages: messages,
	}

	e := validation.Validate(opts)

	ve := s.Version.Validate()
	out := validation.MergeErrors(e, ve)

	return out
}

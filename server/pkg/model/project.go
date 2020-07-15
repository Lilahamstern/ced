package model

import (
	"github.com/lilahamstern/ced/server/pkg/validation"
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type (
	Project struct {
		ID        int64  `json:"id"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}

	CreateProject struct {
		ID          int64             `json:"id"`
		Information CreateInformation `json:"information"`
	}
)

func (s *CreateProject) Validate() url.Values {
	rules := govalidator.MapData{
		"id": []string{"required", "min:9999999", "unique:projects"},
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

	ve := s.Information.Validate()
	out := validation.MergeErrors(e, ve)

	return out
}

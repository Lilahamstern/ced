package model

import (
	"github.com/thedevsaddam/govalidator"
	"net/url"
)

type (
	Project struct {
		ID        int64 `json:"id"`
		CreatedAt string
		UpdatedAt string
	}

	CreateProject struct {
		ID      int64         `json:"id"`
		Version CreateVersion `json:"version"`
	}

	DeleteProject struct {
	}
)

func (s *CreateProject) Validate() url.Values {
	rules := govalidator.MapData{
		"id": []string{"required", "min:9999999"},
	}

	messages := govalidator.MapData{
		"id": []string{"min:The id field cannot be less then 8 numbers"},
	}

	opts := govalidator.Options{
		Data:     s,
		Rules:    rules,
		Messages: messages,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()

	ve := s.Version.Validate()
	if len(ve) != 0 {
		for k, v := range ve {
			for _, vv := range v {
				e.Add(k, vv)
			}
		}
	}

	if len(e) == 0 {
		return nil
	}

	return e
}

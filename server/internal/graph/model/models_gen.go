// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateProjectInput struct {
	ID int64 `json:"id"`
}

type CreateVersionInput struct {
	ProjectID     int64  `json:"projectId"`
	InformationID string `json:"informationId"`
}

type DeleteProjectInput struct {
	ID int64 `json:"id"`
}

type DeleteVersionInput struct {
	ID string `json:"id"`
}

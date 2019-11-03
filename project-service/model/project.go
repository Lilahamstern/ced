package model

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	ServedBy string `json:"served_by"`
}

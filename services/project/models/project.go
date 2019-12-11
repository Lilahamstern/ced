package models

// Project struct
type Project struct {
	ID     string `json:"id"`
	Client string `json:"client"`
	Sector string `json:"sector"`
	Name   string `json:"name"`
	Model  string `json:"model"`
	Desc   string `json:"desc"`
	Phase  string `json:"phase"`
}

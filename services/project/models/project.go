package models

// Project struct
type Project struct {
	ID     string `json:"id"`
	Client string `json:"client"`
	Sector string `json:"sector"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Base
}

// TableName set database table name
func (Project) TableName() string {
	return "Projects"
}

package models

// Project struct
type Project struct {
	PID    string `json:"pid"`
	UID    string `json:"uid"`
	Client string `json:"client"`
	Sector string `json:"sector"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Base
}

// TableName set database table name
func (Project) TableName() string {
	return "Projects"
}

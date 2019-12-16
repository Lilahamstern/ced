package models

// Component struct
type Component struct {
	PID      string  `json:"pid,omitempty"`
	Name     string  `json:"name"`
	Profile  string  `json:"profile"`
	Material string  `json:"material"`
	CO       float32 `json:"co"`
	Type     string  `json:"type"`
	Base
}

// TableName set database table name
func (Component) TableName() string {
	return "Components"
}

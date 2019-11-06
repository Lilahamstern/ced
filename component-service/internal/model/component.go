package model

type Component struct {
	ProjectID string  `json:"project_id"`
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Profile   string  `json:"profile"`
	Material  string  `json:"material"`
	Class     string  `json:"class"`
	Co2       float64 `json:"co_2"`
	Version   int     `json:"version"`
}

type Components []Component

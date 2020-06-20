package model

type Version struct {
	ID            string `json:"id"`
	ProjectId     int64  `json:"projectId"`
	InformationId string `json:"informationId"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

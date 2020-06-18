package model

type Project struct {
	ID        string   `json:"id"`
	Version   *Version `json:"version"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

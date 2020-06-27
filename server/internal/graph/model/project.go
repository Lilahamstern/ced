package model

type Project struct {
	ID        int64    `json:"id"`
	Version   *Version `json:"version"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}
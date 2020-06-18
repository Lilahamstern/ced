package graph

type Project struct {
	ID        string   `json:"id"`
	VersionId *Version `json:"version"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
}

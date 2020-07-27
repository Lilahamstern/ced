package response

type Project struct {
	ID        int64     `json:"id"`
	Versions  []Version `json:"versions"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

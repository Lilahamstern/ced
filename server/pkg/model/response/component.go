package response

type Component struct {
	ID        int64   `json:"id"`
	Name      string  `json:"name"`
	Profile   string  `json:"profile"`
	Material  string  `json:"material"`
	Co        float64 `json:"co"`
	Level     int64   `json:"level"`
	Type      string  `json:"type"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

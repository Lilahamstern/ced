package response

type Version struct {
	ID          string  `json:"id,omitempty"`
	OrderID     int64   `json:"order_id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description *string `json:"desc,omitempty"`
	Manager     string  `json:"manager,omitempty"`
	Client      string  `json:"client,omitempty"`
	Sector      string  `json:"sector,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
}

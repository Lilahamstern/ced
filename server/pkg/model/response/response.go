package response

// Response : Response model
type Response struct {
	Status  string      `json:"status"`
	ID      interface{} `json:"error_id,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

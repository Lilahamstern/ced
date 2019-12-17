package models

// DataResponse struct
type DataResponse struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

// MessageResponse struct
type MessageResponse struct {
	Status   string   `json:"status"`
	Message  string   `json:"message,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

package model

type Message struct {
	Message string `json:"message"`
}

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

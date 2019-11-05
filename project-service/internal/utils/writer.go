package utils

import (
	"encoding/json"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model/response"
	"net/http"
	"strconv"
)

func WriteJsonResponse(w http.ResponseWriter, status int, res interface{}) {
	data, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func WriteJsonMessage(w http.ResponseWriter, status int, msg string) {
	data, _ := json.Marshal(response.Message{Message: msg})
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

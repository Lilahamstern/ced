package utils

import (
	"encoding/json"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model/response"
	"net/http"
	"strconv"
)

func WriteJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	res := model.Response{
		Status: status,
		Data:   data,
	}
	write(w, res)
}

func WriteJsonMessage(w http.ResponseWriter, status int, msg string) {
	res := model.Response{
		Status: status,
		Data:   model.Message{Message: msg},
	}
	write(w, res)
}

func write(w http.ResponseWriter, res model.Response) {
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(res.Status)
	_, _ = w.Write(data)
}

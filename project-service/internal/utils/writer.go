package utils

import (
	"encoding/json"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model/response"
	"net/http"
	"strconv"
)

func WriteJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	res := model.Response{
		Status: getStatus(status),
		Data:   data,
	}
	write(w, res, status)
}

func WriteJsonMessage(w http.ResponseWriter, status int, msg string) {
	res := model.Response{
		Status: getStatus(status),
		Data:   model.Message{Message: msg},
	}
	write(w, res, status)
}

func write(w http.ResponseWriter, res model.Response, status int) {
	data, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func getStatus(status int) string {

	if status >= 500 {
		return "I just fucked it up"
	}

	if status >= 400 {
		return "You just fucked something up"
	}

	if status >= 300 {
		return "Please leave me alone, i will just say no"
	}

	if status >= 200 {
		return "Succeed to serve your fucking data"
	}

	if status >= 100 {
		return "Can you wait a bit."
	}

	return ""
}

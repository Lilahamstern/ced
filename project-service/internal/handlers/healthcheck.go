package handlers

import (
	"encoding/json"
	"net/http"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	dbUp := DBClient.Check()

	if !dbUp {
		data, _ := json.Marshal(healthCheckResponse{Status: "DOWN"})
		writeJsonResponse(w, http.StatusServiceUnavailable, data)
		return
	}

	data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
	writeJsonResponse(w, http.StatusOK, data)
}

package handlers

import (
	"encoding/json"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"net/http"
)

func CreateProject(w http.ResponseWriter, r *http.Request) {

	var req model.Project

	decoder := json.NewDecoder(r.Body)

	_ = decoder.Decode(&req)

	project, err := DBClient.CreateProject(req)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(project)

	writeJsonResponse(w, http.StatusCreated, data)

}

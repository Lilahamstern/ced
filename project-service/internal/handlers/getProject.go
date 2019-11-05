package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetProject(w http.ResponseWriter, r *http.Request) {

	var projectId = mux.Vars(r)["projectId"]

	project, err := DBClient.QueryProject(projectId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(project)

	writeJsonResponse(w, http.StatusCreated, data)

}

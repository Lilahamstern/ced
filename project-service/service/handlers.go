package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lilahamstern/bec/project-service/service/dbclient"
	"net/http"
	"strconv"
)

var DBClient dbclient.IBoltClient

func GetProject(w http.ResponseWriter, r *http.Request) {

	var projectId = mux.Vars(r)["projectId"]

	project, err := DBClient.QueryProject(projectId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(project)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

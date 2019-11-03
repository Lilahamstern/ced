package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lilahamstern/bec-microservices/project-service/service/dbclient"
	"net"
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

	project.ServedBy = getIP()

	data, _ := json.Marshal(project)

	writeJsonResponse(w, http.StatusCreated, data)

}

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

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	_, _ = w.Write(data)
}

func getIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}

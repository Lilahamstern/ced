package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/project-service/internal/database"
	"github.com/lilahamstern/bec-microservices/project-service/internal/utils"
	"net/http"
)

type healthCheckResponse struct {
	Status string `json:"status"`
}

func HealthCheck(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)
	dbUp := db.Check()

	if !dbUp {
		data, _ := json.Marshal(healthCheckResponse{Status: "DOWN"})
		utils.WriteJsonResponse(c.Writer, http.StatusServiceUnavailable, data)
		return
	}

	data, _ := json.Marshal(healthCheckResponse{Status: "UP"})
	utils.WriteJsonResponse(c.Writer, http.StatusOK, data)
}

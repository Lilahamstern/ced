package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/component-service/internal/database"
	"github.com/lilahamstern/bec-microservices/component-service/internal/utils"
	"net/http"
)

func GetComponents(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var projectId = c.Param("projectId")

	components, err := db.QueryComponents(projectId)

	if err != nil {
		if err.Error() == "record not found" {
			utils.WriteJsonMessage(c.Writer, http.StatusNotFound, "components not found")
			return
		}
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJsonResponse(c.Writer, http.StatusOK, components)

}

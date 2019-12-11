package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/project-service/internal/database"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"github.com/lilahamstern/bec-microservices/project-service/internal/utils"
	"net/http"
)

func DeleteProject(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var projectId = c.Param("projectId")

	if exists := db.ProjectExists(model.Project{
		Id: projectId,
	}); exists == nil {
		utils.WriteJsonMessage(c.Writer, http.StatusNotFound, "Project could not be found")
		return
	}

	err := db.DeleteProject(projectId)

	if err != nil {
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJsonMessage(c.Writer, http.StatusAccepted, "project deleted")
}

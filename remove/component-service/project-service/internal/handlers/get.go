package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/project-service/internal/database"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"github.com/lilahamstern/bec-microservices/project-service/internal/utils"
	"net/http"
)

func GetProject(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var projectId = c.Param("projectId")

	project, err := db.QueryProject(projectId)

	if err != nil {
		if err.Error() == "record not found" {
			utils.WriteJsonMessage(c.Writer, http.StatusNotFound, "project not found")
			return
		}
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJsonResponse(c.Writer, http.StatusOK, project)
}

func GetAllProjects(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var projects model.Projects
	var err error

	search := c.Query("search")
	if search == "" {
		projects, err = db.QueryAllProjects()
	}

	if len(search) > 0 {
		projects, err = db.SearchProjects(search)
	}

	if err != nil {
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	if len(projects) < 1 {
		utils.WriteJsonMessage(c.Writer, http.StatusNotFound, "projects not found")
		return
	}

	utils.WriteJsonResponse(c.Writer, http.StatusOK, projects)
}
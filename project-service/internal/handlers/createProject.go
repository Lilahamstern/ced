package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/project-service/internal/database"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"github.com/lilahamstern/bec-microservices/project-service/internal/utils"
	"net/http"
)

func CreateProject(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var req model.Project

	if e := c.BindJSON(&req); e != nil {
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, e.Error())
	}

	project, err := db.CreateProject(req)

	if err != nil {
		utils.WriteJsonMessage(c.Writer, http.StatusNotFound, err.Error())
		return
	}

	utils.WriteJsonResponse(c.Writer, http.StatusCreated, project)

}

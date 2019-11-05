package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/project-service/internal/database"
	"github.com/lilahamstern/bec-microservices/project-service/internal/utils"
	"net/http"
)

func GetProject(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var projectId = c.Param("projectId")

	project, err := db.QueryProject(projectId)

	fmt.Println(err)

	if err != nil {
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, err.Error())
		return
	}

	utils.WriteJsonResponse(c.Writer, http.StatusCreated, project)

}

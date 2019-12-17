package handlers

import (
	"fmt"
	"net/http"
	"project/database"
	"project/models"
	"project/route"

	"github.com/gin-gonic/gin"
)

func init() {
	route.NewEndpoint("CreateProject", "POST", "/", create).Add()
	route.NewEndpoint("GetAll", "GET", "/", getAll).Add()
}

func create(c *gin.Context) {
	db := c.MustGet("db").(database.DBClient)

	var req models.Project

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, models.MessageResponse{Status: "Error", Message: "Error while fetching data: " + err.Error()})
	}

	project, err := db.CreateProject(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.MessageResponse{Status: "Error", Message: "Error while creating project: " + err.Error()})
		return
	}

	c.JSON(200, models.DataResponse{Status: "Success", Data: project})

}

func getAll(c *gin.Context) {
	db := c.MustGet("db").(database.DBClient)

	var projects []models.Project
	var err error

	limit := c.Query("limit")
	searchQuery := c.Query("search")

	if len(searchQuery) > 1 {
		if len(searchQuery) < 3 {
			c.JSON(http.StatusInternalServerError, models.MessageResponse{Status: "Error", Message: "Search query needs a length of 3."})
			return
		}
		projects, err = db.SearchProjects(searchQuery, limit)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.MessageResponse{Status: "Error", Messages: "Error while searching project: " + err.Error()})
			return
		}
	}

	if searchQuery == "" {
		projects, err = db.QueryAllProjects(limit)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	c.JSON(200, projects)
}

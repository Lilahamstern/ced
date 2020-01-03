package handlers

import (
	"fmt"
	"project/database"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var project database.Project

<<<<<<< HEAD
	if err := c.BindJSON(&project); err != nil {
		fmt.Println(err)
	}

	writeMessage(c, 201, "Project created")
}

func getAll(c *gin.Context) {
	// Why should I implement this shit
	writeMessage(c, 200, "Not implemented")
=======
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
>>>>>>> 90d0f062f00a7016bddf292fe52152c96f40a496
}

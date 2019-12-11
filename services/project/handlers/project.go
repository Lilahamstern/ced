package handlers

import (
	"fmt"
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
		fmt.Println(err)
	}

	project, err := db.CreateProject(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, project)

}

func getAll(c *gin.Context) {
	db := c.MustGet("db").(database.DBClient)

	projects, err := db.QueryAllProjects(c.Query("limit"))
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(200, projects)
}

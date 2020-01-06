package handlers

import (
	"fmt"
	"project/database"

	"github.com/gin-gonic/gin"
)

func create(c *gin.Context) {
	var project database.Project

	if err := c.BindJSON(&project); err != nil {
		fmt.Println(err)
	}

	writeMessage(c, 201, "Project created")
}

func getAll(c *gin.Context) {
	// Why should I implement this shit
	writeMessage(c, 200, "Not implemented")
}

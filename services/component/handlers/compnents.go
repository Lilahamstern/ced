package handlers

import (
	"component/models"
	"component/route"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	route.NewEndpoint("CreateComponents", "POST", "/", create).Add()
}

type componentCreate struct {
	PID       string             `json:"pid"`
	Area      float32            `json:"area"`
	Component []models.Component `json:"comp"`
}

type componentCreateRes struct {
	PID   string `json:"pid"`
	Total int    `json:"total"`
}

func create(c *gin.Context) {
	//db := c.MustGet("db").(database.DBClient)

	fmt.Println(c.PostForm("pid"))
	var req componentCreate

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(404, err)
		return
	}

	res := componentCreateRes{PID: req.PID, Total: len(req.Component)}
	c.JSON(200, res)

}

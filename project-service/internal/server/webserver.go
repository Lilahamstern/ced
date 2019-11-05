package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StartWebService(port string, engine *gin.Engine) {
	log.Println("Starting HTTP server at " + port)
	err := engine.Run(":" + port)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

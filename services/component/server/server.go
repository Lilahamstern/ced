package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

// StartHTTPServer starts http server
func StartHTTPServer(port string, engine *gin.Engine) {
	fmt.Println("Starting HTTP server on port:", port)
	err := engine.Run(":" + port)
	log.Println("error occured strting HTTP server:", err)
}

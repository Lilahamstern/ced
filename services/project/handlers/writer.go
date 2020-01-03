package handlers

import "github.com/gin-gonic/gin"

func writeError(c *gin.Context, status int, errors []error) {
	if len(errors) > 1 {
		c.JSON(status, gin.H{"status": status, "errors": errors})
		return
	}

	c.JSON(status, gin.H{"status": status, "error": errors[0]})
}

func writeData(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data, "status": status})
}

func writeMessage(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
	})
}

package controller

import (
	"github.com/gofiber/fiber"
	"log"
)

func respondJSON(c *fiber.Ctx, status int, payload interface{}) {
	if err := c.Status(status).JSON(payload); err != nil {
		log.Fatalln(err)
	}
}

func respondData(c *fiber.Ctx, status int, payload interface{}) {
	respondJSON(c, status, map[string]interface{}{"data": payload})
}

func respondError(c *fiber.Ctx, status int, payload interface{}) {
	respondJSON(c, status, map[string]interface{}{"errors": payload})
}

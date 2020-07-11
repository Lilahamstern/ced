package server

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/controller"
)

func registerV1Routes(api *fiber.Group, controller *controller.Controller) {
	v1 := api.Group("/v1")
	{
		v1.Post("/projects", controller.CreateProject)
	}
}

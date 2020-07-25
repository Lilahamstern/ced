package space

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/repository"
)

type Handler struct {
	Repos *repository.Repositories
}

func (h *Handler) Routes(space *gin.RouterGroup) {
	space.POST("/projects", h.handleProjectCreate())
	space.GET("/projects", h.handleProjectGetAll())
	space.GET("/projects/:id", h.handleProjectGet())
}

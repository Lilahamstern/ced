package space

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/repository"
)

// Handler : Space handler struct
type Handler struct {
	Repos *repository.Repositories
}

// Routes : Register space routes
func (h *Handler) Routes(gin *gin.Engine) {
	space := gin.Group("/space")

	space.POST("/projects", h.handleProjectCreate())
	space.GET("/projects", h.handleProjectGetAll())
	space.GET("/projects/:id", h.handleProjectGet())
}

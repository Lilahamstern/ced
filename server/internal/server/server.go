package server

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository"
	"github.com/lilahamstern/ced/server/internal/server/handler/middleware"
	"github.com/lilahamstern/ced/server/internal/server/handler/space"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Server struct {
	db     *sqlx.DB
	router *gin.Engine
}

func NewServer(db *sqlx.DB) *Server {
	r := gin.New()
	s := &Server{db, r}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())
	s.router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	s.router.Use(middleware.ErrorHandler())

	spaceHandler := &space.Handler{
		Repos: repository.New(s.db),
	}

	api := s.router.Group("/api")

	spaceHandler.Routes(api.Group("/space"))

}

func (s *Server) Start() {
	port := viper.GetString("port")
	log.Println("Starting server")
	log.Printf("Server started! Visit http://localhost:%v/ for documentation!", port)
	if err := s.router.Run(port); err != nil {
		log.Fatalln(err)
	}
}

package server

import (
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository"
	"github.com/lilahamstern/ced/server/internal/server/handler/middleware"
	"github.com/lilahamstern/ced/server/internal/server/handler/space"
	"github.com/lilahamstern/ced/server/pkg/config"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	db     *sqlx.DB
	router *gin.Engine
}

func NewServer(db *sqlx.DB) *Server {
	r := gin.New()
	s := &Server{db, r}
	s.middleware()
	s.routes()

	return s
}

func (s *Server) middleware() {
	s.router.Use(gin.Recovery())
	s.router.Use(gin.Logger())
	s.router.Use(middleware.CORS())
}

func (s *Server) routes() {
	s.router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	s.router.Use(middleware.ErrorHandler())

	spaceHandler := &space.Handler{
		Repos: repository.New(s.db),
	}

	spaceHandler.Routes(s.router)

}

func (s *Server) Start() {
	port := config.GetPort()
	log.Println("Starting server")
	log.Printf("Server started! Visit http://localhost:%v/ for documentation!", port)
	if err := s.router.Run(port); err != nil {
		log.Fatalln(err)
	}
}

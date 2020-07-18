package server

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/lilahamstern/ced/server/internal/repository"
	"github.com/lilahamstern/ced/server/internal/server/handler"
	"github.com/lilahamstern/ced/server/internal/server/handler/space"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	db  *sqlx.DB
	app *fiber.App
}

func NewServer(db *sqlx.DB) *Server {
	app := fiber.New(&fiber.Settings{
		ErrorHandler:          handler.ErrorHandler,
		StrictRouting:         true,
		DisableStartupMessage: false,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
	})
	s := &Server{db, app}
	s.routes()

	return s
}

func (s *Server) routes() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger(middleware.LoggerConfig{
		Format:     "${time} ${method} ${path} - ${ip} - ${status} - ${latency}\n",
		TimeFormat: "15:04:05",
	}))

	spaceHandler := &space.Handler{
		Repos: repository.New(s.db),
	}

	api := s.app.Group("/api")
	{
		spaceHandler.Routes(api.Group("/space"))
	}
}

func (s *Server) Start(port int) {
	go func() {
		log.Println("Starting server")
		log.Printf("Server started! Visit http://localhost:%v/ for documentation!", port)
		if err := s.app.Listen(port); err != nil {
			log.Fatal(err)
		}
	}()
}

func (s *Server) SafeShutDown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal is received
	<-interruptChan

	if err := s.app.Shutdown(); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down server...")
	os.Exit(0)
}

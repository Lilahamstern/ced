package server

import (
	"github.com/gofiber/fiber"
	"github.com/lilahamstern/ced/server/internal/controller"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	app  *fiber.App
	port string
}

func New(port string, controller *controller.Controller) *Server {
	app := fiber.New(&fiber.Settings{
		ServerHeader:          "CED",
		StrictRouting:         true,
		DisableStartupMessage: false,
		ReadTimeout:           5 * time.Second,
		WriteTimeout:          5 * time.Second,
	})

	api := app.Group("/api")

	registerV1Routes(api, controller)

	return &Server{
		app,
		port,
	}
}

func (s *Server) Start() {
	go func() {
		log.Println("Starting server")
		log.Printf("Server started! Visit http://localhost:%s/ for GrahQL Playground", s.port)
		if err := s.app.Listen(s.port); err != nil {
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

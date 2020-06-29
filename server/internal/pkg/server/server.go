package server

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/ced/server/internal/pkg/server/handler"
	"github.com/lilahamstern/ced/server/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	srv  *http.Server
	port string
}

func NewHttpServer(conf *config.Config) *Server {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.POST("/query", handler.GraphQLHandler())
	router.GET("/", handler.PlaygroundHandler())

	srv := &http.Server{
		Addr:           ":" + conf.Port,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &Server{
		srv,
		conf.Port,
	}
}

func (s *Server) Start() {

	go func() {
		log.Println("Starting server")
		log.Printf("Server started! Visit http://localhost:%s/ for GrahQL Playground", s.port)
		if err := s.srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
}

func (s *Server) SafeShutDown() {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal is received
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.srv.Shutdown(ctx)

	log.Println("Shutting down server!")
	os.Exit(0)
}

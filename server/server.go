package main

import (
	"context"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/ced/server/pkg/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/lilahamstern/ced/server/internal/pkg/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")

	config.LoadConfig()

	database.InitDB()
	database.Migrate()

	router := gin.Default()
	router.POST("/query", GraphQLHandler())
	router.GET("/", PlaygroundHandler())

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		log.Println("Starting server")
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	calmShutdown(srv)
}

func calmShutdown(srv *http.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until signal is received
	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)

	log.Println("Shutting down server!")
	os.Exit(0)
}

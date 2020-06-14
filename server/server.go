package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/lilahamstern/ced/server/internal/pkg/handler"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}

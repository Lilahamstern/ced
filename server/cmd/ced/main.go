package main

import (
	"github.com/lilahamstern/ced/server/internal/pkg/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"os"
)

func main() {
	config.LoadConfig()
	port := os.Getenv("PORT")

	srv := server.SetupHttpServer(port)

	srv.CreateConnectionServices()
	srv.Start()

	srv.SafeShutDown()
}

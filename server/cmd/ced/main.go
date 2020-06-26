package main

import (
	"fmt"
	"github.com/lilahamstern/ced/server/internal/pkg/server"
	"github.com/lilahamstern/ced/server/pkg/config"
	"os"
)

func main() {
	conf := config.NewConfig()
	fmt.Println(conf)
	port := os.Getenv("PORT")

	srv := server.SetupHttpServer(port)

	srv.CreateConnectionServices()
	srv.Start()

	srv.SafeShutDown()
}

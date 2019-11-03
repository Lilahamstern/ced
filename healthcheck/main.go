package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	port := flag.String("port", "80", "port on localhost to service")
	flag.Parse()

	res, err := http.Get("http://localhost:" + *port + "/health")

	if err != nil || res.StatusCode != 200 {
		os.Exit(1)
	}

	os.Exit(0)
}

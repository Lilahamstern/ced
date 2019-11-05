package service

import (
	"log"
	"net/http"
)

func StartWebService(port string) {
	log.Println("Starting HTTP service at " + port)
	err := http.ListenAndServe(":"+port, NewRouter()) // Goroutine will block here
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}

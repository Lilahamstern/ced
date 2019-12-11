package services

import (
	"fmt"
	"net/http"
)

func CheckProjectExists(projectId string) bool {
	res, err := http.Get("http://bec.service.project:5050/projects/" + projectId)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	if res.StatusCode == 200 {
		return true
	}

	return false
}

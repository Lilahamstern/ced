package handlers

import (
	"project/config"

	"github.com/gin-gonic/gin"
)

type infoResponse struct {
	General generalInfo    `json:"general"`
	Service config.Service `json:"service"`
}

type generalInfo struct {
	IP        string `json:"ip"`
	Endpoints int    `json:"endpoints"`
}

func info(c *gin.Context) {
	cfg := c.MustGet("cfg").(config.Service)
	res := infoResponse{
		General: generalInfo{
			IP:        c.ClientIP(),
			Endpoints: len(c.HandlerNames()),
		},
		Service: cfg,
	}
	writeData(c, 200, res)
}

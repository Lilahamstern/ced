package handlers

import (
	"template/config"
	"template/route"

	"github.com/gin-gonic/gin"
)

func init() {
	route.NewEndpoint("GET", "/info", info).Add()
}

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
			IP:        "10.0.0.1",
			Endpoints: len(c.HandlerNames()) - 4,
		},
		Service: cfg,
	}
	c.JSON(200, res)
}

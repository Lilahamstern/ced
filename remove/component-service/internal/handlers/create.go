package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/lilahamstern/bec-microservices/component-service/internal/database"
	"github.com/lilahamstern/bec-microservices/component-service/internal/model"
	"github.com/lilahamstern/bec-microservices/component-service/internal/services"
	"github.com/lilahamstern/bec-microservices/component-service/internal/utils"
	"net/http"
)

type requestCreate struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Profile  string  `json:"profile"`
	Material string  `json:"material"`
	Class    string  `json:"class"`
	Co2      float64 `json:"co_2"`
	Phase    string  `json:"version"`
}

func CreateComponents(c *gin.Context) {
	db := c.MustGet("db").(database.IClient)

	var projectId = c.Param("projectId")

	if !services.CheckProjectExists(projectId) {
		utils.WriteJsonMessage(c.Writer, http.StatusNotFound, "project not found")
		return
	}

	var req []requestCreate

	if e := c.BindJSON(&req); e != nil {
		utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, e.Error())
		return
	}

	for _, component := range req {
		data := model.Component{
			ProjectID: projectId,
			OId:       component.Id,
			Name:      component.Name,
			Profile:   component.Profile,
			Material:  component.Material,
			Class:     component.Class,
			Co2:       component.Co2,
			Phase:     component.Phase,
		}

		err := db.CreateComponent(data)

		if err != nil {
			utils.WriteJsonMessage(c.Writer, http.StatusInternalServerError, err.Error())
			return
		}
	}

	utils.WriteJsonMessage(c.Writer, http.StatusCreated, "components added")

}

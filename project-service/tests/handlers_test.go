package tests

import (
	"encoding/json"
	"fmt"
	"github.com/lilahamstern/bec-microservices/project-service/internal/handlers"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"github.com/lilahamstern/bec-microservices/project-service/internal/service"
	"github.com/lilahamstern/bec-microservices/project-service/internal/service/dbclient"
	. "github.com/smartystreets/goconvey/convey"
	"net/http/httptest"
	"testing"
)

func TestGetProjectWrongPath(t *testing.T) {
	Convey("Given a HTTP request for /invalid/213", t, func() {
		req := httptest.NewRequest("GET", "/invalid/213", nil)
		resp := httptest.NewRecorder()

		Convey("When they request is handled by the Router", func() {
			service.NewRouter().ServeHTTP(resp, req)

			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

func TestGetProject(t *testing.T) {
	mockRepo := &dbclient.MockMongoClient{}

	mockRepo.On("QueryProject", "123").Return(model.Project{
		Id:   "123",
		Name: "Project_123",
	}, nil)

	mockRepo.On("QueryProject", "456").Return(model.Project{}, fmt.Errorf("could not find project: 456"))

	handlers.DBClient = mockRepo

	Convey("Given a HTTP request for /projects/123", t, func() {
		req := httptest.NewRequest("GET", "/projects/123", nil)
		resp := httptest.NewRecorder()
		Convey("When the request is handled by the Router", func() {
			service.NewRouter().ServeHTTP(resp, req)
			Convey("Then the response should be a 200", func() {
				So(resp.Code, ShouldEqual, 200)
				project := model.Project{}
				_ = json.Unmarshal(resp.Body.Bytes(), &project)
				So(project.Id, ShouldEqual, "123")
				So(project.Name, ShouldEqual, "Project_123")
			})
		})
	})

	Convey("Given a HTTP request for /projects/456", t, func() {
		req := httptest.NewRequest("GET", "/projects/456", nil)
		resp := httptest.NewRecorder()
		Convey("When the request is handled by the Router", func() {
			service.NewRouter().ServeHTTP(resp, req)
			Convey("Then the response should be a 404", func() {
				So(resp.Code, ShouldEqual, 404)
			})
		})
	})
}

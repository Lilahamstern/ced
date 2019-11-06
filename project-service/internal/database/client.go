package database

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"log"
	"os"
)

type IClient interface {
	OpenDb()
	Migrate()
	Inject() gin.HandlerFunc
	CreateProject(project model.Project) (model.Project, error)
	QueryProject(projectId string) (model.Project, error)
	QueryAllProjects() ([]model.Project, error)
	DeleteProject(projectId string) error
	Check() bool
}

type Client struct {
	db *gorm.DB
}

func (c *Client) OpenDb() {
	var err error
	c.db, err = gorm.Open("postgres", os.Getenv("DB_CONFIG"))
	if err != nil {
		log.Fatal(err)
	}

	c.db.LogMode(false)
}

func (c *Client) Check() bool {
	return c.db != nil
}

func (c *Client) Inject() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", c)
		ctx.Next()
	}
}

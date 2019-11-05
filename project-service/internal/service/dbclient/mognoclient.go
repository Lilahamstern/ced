package dbclient

import (
	"context"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type IMongoClient interface {
	OpenMongoDb()
	CreateProject(project model.Project) (model.Project, error)
	QueryProject(projectId string) (model.Project, error)
	Check() bool
}

type MongoClient struct {
	mongoDb *mongo.Client
}

func (mc *MongoClient) OpenMongoDb() {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mc.mongoDb, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongodb:27017"))
	if err != nil {
		log.Fatal(err)
	}
}

func (mc *MongoClient) Check() bool {
	return mc.mongoDb != nil
}

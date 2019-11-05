package dbclient

import (
	"context"
	"fmt"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"time"
)

func (mc *MongoClient) QueryProject(projectId string) (model.Project, error) {
	project := model.Project{}

	collection := mc.mongoDb.Database("bec").Collection("projects")

	filter := model.Project{Id: projectId}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(&project)

	fmt.Println(project)

	if err != nil {
		return model.Project{}, err
	}

	return project, nil
}

func (mc *MongoClient) CreateProject(project model.Project) (model.Project, error) {
	collection := mc.mongoDb.Database("bec").Collection("projects")

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, project)

	if err != nil {
		return model.Project{}, err
	}

	id := res.InsertedID

	fmt.Println(id)

	return project, nil
}

package dbclient

import (
	"fmt"
	"github.com/lilahamstern/bec-microservices/project-service/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockMongoClient struct {
	mock.Mock
}

func (m *MockMongoClient) QueryProject(projectId string) (model.Project, error) {
	args := m.Mock.Called(projectId)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *MockMongoClient) Check() bool {
	args := m.Mock.Called()
	fmt.Println(args)
	return args.Get(0).(bool)
}

func (m *MockMongoClient) OpenMongoDb() {
	// Does nothing
}

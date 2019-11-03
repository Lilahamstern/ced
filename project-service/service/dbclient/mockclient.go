package dbclient

import (
	"fmt"
	"github.com/lilahamstern/bec-microservices/project-service/model"
	"github.com/stretchr/testify/mock"
)

type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient) QueryProject(projectId string) (model.Project, error) {
	args := m.Mock.Called(projectId)
	return args.Get(0).(model.Project), args.Error(1)
}

func (m *MockBoltClient) Check() bool {
	args := m.Mock.Called()
	fmt.Println(args)
	return args.Get(0).(bool)
}

func (m *MockBoltClient) OpenBoltDb() {
	// Does nothing
}

func (m *MockBoltClient) Seed() {
	// Does nothing
}

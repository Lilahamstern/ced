package service

import (
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/pkg/domain"
	. "github.com/lilahamstern/ced/server/pkg/error"
	"testing"
	"time"
)

type mockProjectRepository struct {
}

func (m *mockProjectRepository) Delete(id int64) (bool, error) {
	panic("implement me")
}

func (m *mockProjectRepository) ExistsById(id int64) bool {
	if id == 421 {
		return true
	}
	return false
}

func (m *mockProjectRepository) GetAllProjects(projects *[]domain.Project) error {
	return nil
}

func (m *mockProjectRepository) GetProject(project *domain.Project) error {
	return nil
}

func (m *mockProjectRepository) Save(project *domain.Project) error {
	project.UpdatedAt = time.Now().UTC()
	project.CreatedAt = time.Now().UTC()
	return nil
}

var projectServiceTest ProjectService

func TestMain(m *testing.M) {
	repo := &mockProjectRepository{}
	projectServiceTest = NewProjectService(repo)
	m.Run()
}

func TestProjectService_Save_Success(t *testing.T) {
	project := &model.CreateProjectInput{
		ID: 47271,
	}

	_, err := projectServiceTest.Save(project)

	if err != nil {
		t.Errorf("Expected no error but got %s", err.Error())
	}
}

func TestProjectService_Save_Exists(t *testing.T) {
	project := &model.CreateProjectInput{
		ID: 421,
	}

	_, err := projectServiceTest.Save(project)

	notFoundErr := &ProjectNotFound{ID: project.ID}
	if err == nil {
		t.Errorf("Expected error but got %v", err)
	}

	if err.Error() != notFoundErr.Error() {
		t.Errorf("Expected %s but got %s", notFoundErr.Error(), err.Error())
	}

}

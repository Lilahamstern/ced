package service

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/db/domain"
	"github.com/lilahamstern/ced/server/internal/repository"
	. "github.com/lilahamstern/ced/server/pkg/error"
	"github.com/lilahamstern/ced/server/pkg/model"
	"log"
)

type ProjectService interface {
	Save(input *model.CreateProject) (*model.Project, error)
	GetAll() ([]*model.Project, error)
	Get(id int64) (*model.Project, error)
	Delete(id int64) (bool, error)
}

type projectService struct {
	projectRepository repository.ProjectRepository
}

func (s *projectService) Delete(id int64) (bool, error) {
	exists := s.projectRepository.ExistsById(id)
	if !exists {
		return false, &ProjectNotFound{ID: id}
	}

	deleted, err := s.projectRepository.Delete(id)
	if err != nil {
		return false, &InternalServerError{}
	}

	return deleted, nil
}

func (s projectService) Save(input *model.CreateProject) (*model.Project, error) {
	var project = domain.Project{
		ID: input.ID,
	}

	if s.projectRepository.ExistsById(project.ID) {
		return nil, &ProjectNotFound{ID: project.ID}
	}

	err := s.projectRepository.Save(&project)
	if err != nil {
		log.Println(err)
		return nil, &InternalServerError{}
	}

	return project.ToModel(), nil
}

func (s projectService) GetAll() ([]*model.Project, error) {
	var projects []domain.Project
	err := s.projectRepository.GetAllProjects(&projects)
	if err != nil {
		return nil, err
	}

	var res []*model.Project
	for _, project := range projects {
		res = append(res, project.ToModel())
	}

	return res, nil
}

func (s projectService) Get(id int64) (*model.Project, error) {
	project := domain.Project{
		ID: id,
	}
	err := s.projectRepository.GetProject(&project)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ProjectNotFound{ID: id}
		}
		return nil, &InternalServerError{}
	}

	return project.ToModel(), nil
}

func newProject(projectRepository repository.ProjectRepository) ProjectService {
	return &projectService{
		projectRepository,
	}
}

package service

import (
	"database/sql"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/pkg/repository"
	"github.com/lilahamstern/ced/server/pkg/domain"
	. "github.com/lilahamstern/ced/server/pkg/error"
	"log"
)

type ProjectService struct {
	ProjectRepository *repository.ProjectRepository
}

func (s ProjectService) Save(input *model.CreatProjectInput) (*model.Project, error) {
	var project = domain.Project{
		ID: input.ID,
	}

	if s.ProjectRepository.ExistsById(project.ID) {
		return nil, &ProjectNotFound{ID: project.ID}
	}

	err := s.ProjectRepository.Save(&project)
	if err != nil {
		log.Println(err)
		return nil, &InternalServerError{}
	}

	return project.ToGraphModel(), nil
}

func (s ProjectService) GetAll() ([]*model.Project, error) {
	var projects []domain.Project
	err := s.ProjectRepository.GetAllProjects(&projects)
	if err != nil {
		return nil, err
	}
	var res []*model.Project
	for _, version := range projects {
		res = append(res, version.ToGraphModel())
	}

	return res, nil
}

func (s ProjectService) Get(id int64) (*model.Project, error) {
	project := domain.Project{
		ID: id,
	}
	err := s.ProjectRepository.GetProject(&project)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &ProjectNotFound{ID: id}
		}
		return nil, &InternalServerError{}
	}

	return project.ToGraphModel(), nil
}

func NewProjectService(projectRepository *repository.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepository,
	}
}

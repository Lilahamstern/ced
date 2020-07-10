package service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/repository"
	"github.com/lilahamstern/ced/server/pkg/domain"
	. "github.com/lilahamstern/ced/server/pkg/error"
	"log"
	"strconv"
)

type VersionService interface {
	Save(input *model.CreateVersionInput) (*model.Version, error)
	GetByProjectId(id int64) ([]*model.Version, error)
	GetById(id string) (*model.Version, error)
	Delete(input *model.DeleteVersionInput) (bool, error)
}

type versionService struct {
	versionRepository repository.VersionRepository
	projectRepository repository.ProjectRepository
}

func (s versionService) Delete(input *model.DeleteVersionInput) (bool, error) {
	UUID, err := uuid.Parse(input.ID)
	if err != nil {
		log.Fatal(err)
		return false, &InternalServerError{}
	}

	exists := s.versionRepository.ExistsById(UUID)
	if !exists {
		return false, &VersionNotFound{ID: UUID.String()}
	}

	deleted, err := s.versionRepository.Delete(UUID)
	if err != nil {
		return false, &InternalServerError{}
	}

	return deleted, nil
}

func (s versionService) Save(input *model.CreateVersionInput) (*model.Version, error) {
	UUID, err := uuid.Parse(input.InformationID)

	if err != nil {
		log.Fatal(err)
		return nil, &InternalServerError{}
	}

	version := domain.Version{
		ProjectId:     input.ProjectID,
		InformationId: UUID,
	}

	if !s.projectRepository.ExistsById(input.ProjectID) {
		return nil, &VersionNotFound{ID: strconv.FormatInt(input.ProjectID, 10)}
	}

	err = s.versionRepository.Save(&version)
	if err != nil {
		log.Println(err)
		return nil, &InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

func (s versionService) GetByProjectId(id int64) ([]*model.Version, error) {
	var versions []domain.Version
	err := s.versionRepository.GetVersionByProjectId(id, &versions)
	if err != nil {
		return nil, err
	}

	var res []*model.Version
	for _, version := range versions {
		res = append(res, version.ToGraphModel())
	}

	return res, nil
}

func (s versionService) GetById(id string) (*model.Version, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	version := domain.Version{
		ID: parsedId,
	}

	err = s.versionRepository.GetVersionById(version)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &VersionNotFound{ID: id}
		}
		return nil, &InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

func newVersion(versionRepository repository.VersionRepository,
	projectRepository repository.ProjectRepository) VersionService {
	return &versionService{
		versionRepository,
		projectRepository,
	}
}

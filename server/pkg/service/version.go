package service

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/pkg/repository"
	"github.com/lilahamstern/ced/server/pkg/domain"
	. "github.com/lilahamstern/ced/server/pkg/error"
	"log"
	"strconv"
)

type VersionService struct {
	VersionRepository *repository.VersionRepository
	ProjectRepository *repository.ProjectRepository
}

func (s VersionService) Save(input *model.CreateVersionInput) (*model.Version, error) {
	parsedInfoId, err := uuid.Parse(input.InformationID)

	if err != nil {
		log.Fatal(err)
		return nil, &InternalServerError{}
	}

	version := domain.Version{
		ProjectId:     input.ProjectID,
		InformationId: parsedInfoId,
	}

	if !s.ProjectRepository.ExistsById(input.ProjectID) {
		return nil, &VersionNotFound{ID: strconv.FormatInt(input.ProjectID, 10)}
	}

	err = s.VersionRepository.Save(&version)
	if err != nil {
		log.Println(err)
		return nil, &InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

func (s VersionService) GetByProjectId(id int64) ([]*model.Version, error) {
	var versions []domain.Version
	err := s.VersionRepository.GetVersionByProjectId(id, &versions)
	if err != nil {
		return nil, err
	}

	var res []*model.Version
	for _, version := range versions {
		res = append(res, version.ToGraphModel())
	}

	return res, nil
}

func (s VersionService) GetById(id string) (*model.Version, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	version := domain.Version{
		ID: parsedId,
	}

	err = s.VersionRepository.GetVersionById(version)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &VersionNotFound{ID: id}
		}
		return nil, &InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

func NewVersionService(versionRepository *repository.VersionRepository, projectRepository *repository.ProjectRepository) *VersionService {
	return &VersionService{
		versionRepository,
		projectRepository,
	}
}
package service

import (
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/pkg/repository"
	"github.com/lilahamstern/ced/server/pkg/domain"
	. "github.com/lilahamstern/ced/server/pkg/error"
)

type VersionInformationService interface {
	Save(input *model.CreateVersionInformationInput) (*model.VersionInformation, error)
}

type versionInformationService struct {
	versionInformationRepository repository.VersionInformationRepository
}

func (s versionInformationService) Save(input *model.CreateVersionInformationInput) (*model.VersionInformation, error) {
	information := domain.VersionInformation{
		OrderID:     input.OrderID,
		Title:       input.Title,
		Description: input.Description,
		Manager:     input.Manager,
		Client:      input.Client,
		Sector:      input.Sector,
	}
	err := s.versionInformationRepository.Save(&information)
	if err != nil {
		return nil, &InternalServerError{}
	}

	return information.ToGraphModel(), err
}

func newVersionInformation(informationRepository repository.VersionInformationRepository) VersionInformationService {
	return &versionInformationService{
		informationRepository,
	}
}

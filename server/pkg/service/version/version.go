package version

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	"github.com/lilahamstern/ced/server/internal/pkg/errors"
	"github.com/lilahamstern/ced/server/internal/pkg/repository"
	"github.com/lilahamstern/ced/server/pkg/service/project"
	"log"
)

func Save(input *model.CreateVersionInput) (*model.Version, error) {
	parsedInfoId, err := uuid.Parse(input.InformationID)

	if err != nil {
		log.Fatal(err)
		return nil, &errors.InternalServerError{}
	}

	version := repository.Version{
		ProjectId:     input.ProjectID,
		InformationId: parsedInfoId,
	}

	if !repository.ProjectExistsById(input.ProjectID) {
		return nil, &project.NotFound{Id: input.ProjectID}
	}

	err = version.Save()
	if err != nil {
		log.Println(err)
		return nil, &errors.InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

func GetByProjectId(id int64) ([]*model.Version, error) {
	var versions []repository.Version
	err := repository.GetVersionByProjectId(id, &versions)
	if err != nil {
		return nil, err
	}

	var res []*model.Version
	for _, version := range versions {
		res = append(res, version.ToGraphModel())
	}

	return res, nil
}

func GetById(id string) (*model.Version, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	version := repository.Version{
		ID: parsedId,
	}

	err = version.GetVersionById()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFound{Id: id}
		}
		return nil, &errors.InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

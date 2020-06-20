package version

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/lilahamstern/ced/server/internal/graph/model"
	database "github.com/lilahamstern/ced/server/internal/pkg/db/postgres"
	"github.com/lilahamstern/ced/server/internal/pkg/errors"
	"github.com/lilahamstern/ced/server/internal/project"
	"log"
)

func Save(input *model.CreateVersionInput) (*model.Version, error) {
	parsedInfoId, err := uuid.Parse(input.InformationID)

	if err != nil {
		log.Fatal(err)
		return nil, &errors.InternalServerError{}
	}

	version := database.Version{
		ProjectId:     input.ProjectID,
		InformationId: parsedInfoId,
	}

	if !database.ProjectExistsById(input.ProjectID) {
		return nil, &project.NotFound{Id: input.ProjectID}
	}

	version, err = version.Save()
	if err != nil {
		log.Println(err)
		return nil, &errors.InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

func GetById(id string) (*model.Version, error) {
	version, err := database.GetVersionById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &NotFound{Id: id}
		}
		return nil, &errors.InternalServerError{}
	}

	return version.ToGraphModel(), nil
}

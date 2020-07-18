package domain

import (
	"github.com/lilahamstern/ced/server/pkg/model/response"
	"time"
)

// Project : Database model
type Project struct {
	ID        int64 `db:"project_id"`
	Versions  []Version
	Version             // Required to scan struct from db.
	CreatedAt time.Time `db:"project_createdat"`
	UpdatedAt time.Time `db:"project_updatedat"`
}

// ToModel : will map database model of project to request model of project
func (p *Project) ToModel() *response.Project {
	versions := []response.Version{}
	// If len(versions) is greater then 0, map values to versions
	if p.Versions != nil {
		for _, v := range p.Versions {
			versions = append(versions, *v.ToModel())
		}
	}

	if p.Version.ID != nil {
		versions = append(versions, *p.Version.ToModel())
	}

	return &response.Project{
		ID:        p.ID,
		Versions:  versions,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

// ToModel : maps slice of domain project to slice of response project model
func ToModel(projects []Project) []response.Project {
	if projects == nil {
		return []response.Project{}
	}

	var p []response.Project
	for _, v := range projects {
		p = append(p, *v.ToModel())
	}

	return p
}

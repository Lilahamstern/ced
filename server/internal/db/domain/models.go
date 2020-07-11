package domain

import (
	"github.com/lilahamstern/ced/server/pkg/model"
	"time"
)

// Project model
type Project struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToGraphModel will map projectRepository to Graphql model of project and return it as pointer
func (p Project) ToModel() *model.Project {
	return &model.Project{
		ID:        p.ID,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

//// Version model for database
//type Version struct {
//	ID            uuid.UUID `json:"id"`
//	ProjectId     int64     `json:"project_id"`
//	InformationId uuid.UUID `json:"information_id"`
//	CreatedAt     time.Time `json:"created_at"`
//	UpdatedAt     time.Time `json:"updated_at"`
//}
//
//func (v Version) ToGraphModel() *model.Version {
//	return &model.Version{
//		ID:            v.ID.String(),
//		ProjectId:     v.ProjectId,
//		InformationId: v.InformationId.String(),
//		CreatedAt:     v.CreatedAt.String(),
//		UpdatedAt:     v.UpdatedAt.String(),
//	}
//}
//
//type VersionInformation struct {
//	ID          uuid.UUID `json:"id"`
//	OrderID     string    `json:"orderId"`
//	Title       string    `json:"title"`
//	Description *string   `json:"description"`
//	Manager     string    `json:"manager"`
//	Client      string    `json:"client"`
//	Sector      string    `json:"sector"`
//	CreatedAt   time.Time `json:"createdAt"`
//	UpdatedAt   time.Time `json:"updatedAt"`
//}
//
//func (v VersionInformation) ToGraphModel() *model.VersionInformation {
//	return &model.VersionInformation{
//		ID:          v.ID.String(),
//		OrderID:     v.OrderID,
//		Title:       v.Title,
//		Description: v.Description,
//		Manager:     v.Manager,
//		Client:      v.Client,
//		Sector:      v.Sector,
//		CreatedAt:   v.CreatedAt.String(),
//		UpdatedAt:   v.UpdatedAt.String(),
//	}
//}

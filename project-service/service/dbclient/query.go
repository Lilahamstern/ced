package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/lilahamstern/bec/project-service/model"
)

func (bc *BoltClient) QueryProject(projectId string) (model.Project, error) {
	project := model.Project{}

	err := bc.boltDB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("ProjectBucket"))

		projectBytes := b.Get([]byte(projectId))
		if projectBytes == nil {
			return fmt.Errorf("no project found for " + projectId)
		}

		json.Unmarshal(projectBytes, &project)

		return nil
	})

	if err != nil {
		return model.Project{}, err
	}

	return project, nil
}

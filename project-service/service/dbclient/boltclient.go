package dbclient

import (
	"github.com/boltdb/bolt"
	"github.com/lilahamstern/bec/project-service/model"
	"log"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryProject(projectId string) (model.Project, error)
	Seed()
}

type BoltClient struct {
	boltDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boltDB, err = bolt.Open("./project-service/projects.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

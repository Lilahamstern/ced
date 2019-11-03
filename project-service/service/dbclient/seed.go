package dbclient

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/lilahamstern/bec-microservices/project-service/model"
	"strconv"
)

func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedProjects()
}

func (bc *BoltClient) initializeBucket() {
	_ = bc.boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("ProjectBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

func (bc *BoltClient) seedProjects() {
	total := 100
	for i := 0; i < total; i++ {

		key := strconv.Itoa(10000 + i)

		acc := model.Project{
			Id:   key,
			Name: "Project_" + strconv.Itoa(i),
		}

		jsonBytes, _ := json.Marshal(acc)

		_ = bc.boltDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("ProjectBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake projects...\n", total)
}

package dbclient

//
//import (
//	"encoding/json"
//	"fmt"
//	"github.com/boltdb/bolt"
//	"github.com/lilahamstern/bec-microservices/project-service/model"
//	"strconv"
//)
//
//func (mc *MongoClient) Seed() {
//	mc.initializeBucket()
//	mc.seedProjects()
//}
//
////func (mc *MongoClient) initializeBucket() {
////	_ = mc.boltDB.Update(func(tx *bolt.Tx) error {
////		_, err := tx.CreateBucket([]byte("ProjectBucket"))
////		if err != nil {
////			return fmt.Errorf("create bucket failed: %s", err)
////		}
////		return nil
////	})
////}
////
////func (mc *MongoClient) seedProjects() {
////	total := 100
////	for i := 0; i < total; i++ {
////
////		key := strconv.Itoa(10000 + i)
////
////		acc := model.Project{
////			Id:   key,
////			Name: "Project_" + strconv.Itoa(i),
////		}
////
////		jsonBytes, _ := json.Marshal(acc)
////
////		_ = mc.boltDB.Update(func(tx *bolt.Tx) error {
////			b := tx.Bucket([]byte("ProjectBucket"))
////			err := b.Put([]byte(key), jsonBytes)
////			return err
////		})
////	}
////	fmt.Printf("Seeded %v fake projects...\n", total)
////}

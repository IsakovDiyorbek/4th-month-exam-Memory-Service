package mongo

import (
	"context"

	"github.com/Exam4/4th-month-exam-Memory-Service/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageM struct {
	mongo *mongo.Database
	memory storage.MemoryI

}

func SetupMongoDBConnection() (storage.StorageMongo, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		// SetAuth(options.Credential{Username: "dior", Password: "20005"})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	mongo := client.Database("memories")

	form := NewMemoryRepo(mongo)

	return &StorageM{
		mongo: mongo,
		memory: form,
	}, nil
}

func (s *StorageM) Memory() storage.MemoryI {
	if s.memory == nil {
		s.memory = NewMemoryRepo(s.mongo)
	}
	return s.memory
}

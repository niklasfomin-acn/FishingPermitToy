/*
	GPT Experiment zum Untersuchungsaspekt Module / Datenbanken

Code Snippet zur Anbindung von AWS RDS Postgres an das Backend
Version: 4
Bemerkungen:
*/

package storage

import (
	"context"
	"fmt"
	"log"
	"rest-backend/types"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBStorage struct {
	db  *mongo.Database
	uri string
}

func NewMongoDBStorage(uri string, dbName string) (*MongoDBStorage, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	db := client.Database(dbName)
	ms := &MongoDBStorage{
		db:  db,
		uri: uri,
	}

	return ms, nil
}

func (ms *MongoDBStorage) SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error) {
	collection := ms.db.Collection("citizenPermit")
	res, err := collection.InsertOne(context.Background(), cpr)
	if err != nil {
		log.Printf("Error saving citizen permit request: %v", err)
		return nil, err
	}
	log.Println("Citizen Permit Request saved successfully")
	return res.InsertedID, nil
}

func (ms *MongoDBStorage) FetchProcessedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ms *MongoDBStorage) FetchPendingCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ms *MongoDBStorage) FetchApprovedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ms *MongoDBStorage) FetchRejectedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

func (ms *MongoDBStorage) FetchCitizenPermitRequestByID(id string) (types.CitizenPermit, error) {
	return types.CitizenPermit{}, nil
}

func (ms *MongoDBStorage) ApproveCitizenPermitRequest(id string) error {
	return nil
}

func (ms *MongoDBStorage) RejectCitizenPermitRequest(id string) error {
	return nil
}

func (ms *MongoDBStorage) ViewPermitStatus(id string) (types.CitizenPermit, error) {
	return types.CitizenPermit{}, nil
}

func (ms *MongoDBStorage) FetchCitizenPermitRequests() ([]types.CitizenPermit, error) {
	return nil, nil
}

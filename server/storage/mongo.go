package storage

import (
	"context"
	"log"
	"rest-backend/types"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client            *mongo.Client
	permitCollection  *mongo.Collection
	citizenCollection *mongo.Collection
}

func NewMongoStorage(uri string) *MongoStorage {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return &MongoStorage{
		client:            client,
		permitCollection:  client.Database("permits").Collection("permit"),
		citizenCollection: client.Database("citizens").Collection("citizen"),
	}
}

func (ms *MongoStorage) SavePermitRequest(pr types.Permit) (interface{}, error) {
	insertRequest, err := ms.permitCollection.InsertOne(context.Background(), pr)
	if err != nil {
		return nil, err
	}
	log.Printf("Permit request saved: %+v\n", pr)
	return insertRequest.InsertedID, nil
}

func (ms *MongoStorage) SaveCitizenRequest(cd types.Citizen) (interface{}, error) {
	// cdCopy := bson.M{
	// 	"permitID": permitID,
	// }

	insertCitizen, err := ms.citizenCollection.InsertOne(context.Background(), cd)
	if err != nil {
		return nil, err
	}
	log.Printf("Citizen request saved: %+v\n", cd)
	return insertCitizen.InsertedID, nil
}

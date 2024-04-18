package storage

import (
	"context"
	"log"
	"rest-backend/types"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client *mongo.Client
	//permitCollection        *mongo.Collection
	//citizenCollection       *mongo.Collection
	citizenPermitCollection *mongo.Collection
}

func NewMongoStorage(uri string) *MongoStorage {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	return &MongoStorage{
		client: client,
		//permitCollection:        client.Database("permits").Collection("permit"),
		//citizenCollection:       client.Database("citizens").Collection("citizen"),
		citizenPermitCollection: client.Database("citizenPermits").Collection("citizenPermit"),
	}
}

func (ms *MongoStorage) SaveCitizenPermitRequest(cpr types.CitizenPermit) (interface{}, error) {
	insertCitizenPermit, err := ms.citizenPermitCollection.InsertOne(context.Background(), cpr)
	if err != nil {
		return nil, err
	}
	log.Printf("Citizen permit request saved: %+v\n", cpr)
	return insertCitizenPermit.InsertedID, nil
}

func (ms *MongoStorage) FetchCitizenPermitRequests() ([]types.CitizenPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := ms.citizenPermitCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []types.CitizenPermit
	for cursor.Next(ctx) {
		var result types.CitizenPermit
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *MongoStorage) FetchProcessedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := ms.citizenPermitCollection.Find(ctx, bson.M{"permitstate": bson.M{"$in": []string{"approved", "rejected"}}})
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []types.CitizenPermit
	for cursor.Next(ctx) {
		var result types.CitizenPermit
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *MongoStorage) FetchPendingCitizenPermitRequests() ([]types.CitizenPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := ms.citizenPermitCollection.Find(ctx, bson.M{"permitstate": "pending"})
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []types.CitizenPermit
	for cursor.Next(ctx) {
		var result types.CitizenPermit
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *MongoStorage) FetchApprovedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := ms.citizenPermitCollection.Find(ctx, bson.M{"permitstate": "approved"})
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []types.CitizenPermit
	for cursor.Next(ctx) {
		var result types.CitizenPermit
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *MongoStorage) FetchRejectedCitizenPermitRequests() ([]types.CitizenPermit, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor, err := ms.citizenPermitCollection.Find(ctx, bson.M{"permitstate": "rejected"})
	if err != nil {
		log.Printf("Error fetching processed citizen permit requests from database: %v\n", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []types.CitizenPermit
	for cursor.Next(ctx) {
		var result types.CitizenPermit
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (ms *MongoStorage) FetchCitizenPermitRequestByID(id string) (types.CitizenPermit, error) {
	// Fetch by Passport Number as unique identifier
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor := ms.citizenPermitCollection.FindOne(ctx, bson.M{"passportnumber": id})

	var result types.CitizenPermit
	err := cursor.Decode(&result)
	if err != nil {
		log.Printf("Error decoding document into citizen permit request: %v\n", err)
		return types.CitizenPermit{}, err
	}

	return result, nil
}

func (ms *MongoStorage) ApproveCitizenPermitRequest(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor := ms.citizenPermitCollection.FindOneAndUpdate(ctx, bson.M{"passportnumber": id}, bson.M{"$set": bson.M{"permitstate": "approved"}})

	var result types.CitizenPermit
	err := cursor.Decode(&result)
	if err != nil {
		log.Printf("Error changing permit status: %v\n", err)
		return err
	}
	return nil
}

func (ms *MongoStorage) RejectCitizenPermitRequest(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cursor := ms.citizenPermitCollection.FindOneAndUpdate(ctx, bson.M{"passportnumber": id}, bson.M{"$set": bson.M{"permitstate": "rejected"}})

	var result types.CitizenPermit
	err := cursor.Decode(&result)
	if err != nil {
		log.Printf("Error changing permit status: %v\n", err)
		return err
	}
	return nil
}

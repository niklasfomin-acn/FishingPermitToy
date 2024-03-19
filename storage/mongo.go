package storage

import "rest-backend/types"

type MongoStorage struct {
	//TODO: Add mongoDB client here
}

func NewMongoStorage() *MongoStorage {
	return &MongoStorage{}
}

func (s *MongoStorage) Get(id int) *types.User {
	return &types.User{
		ID:   id,
		Name: "Test User",
	}
}

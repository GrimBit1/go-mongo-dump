package utils

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func UsersColl() (*mongo.Collection, error) {
	conn, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	db := conn.Database("test")
	usersColl := db.Collection("users")
	return usersColl, nil
}

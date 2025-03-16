package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertOne(users []User) {
	conn, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	db := conn.Database("test")
	usersColl := db.Collection("users")
	for _, user := range users {
		_, err := usersColl.InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Println(err)
			return
		}

		// fmt.Println(res.Acknowledged)
	}
}

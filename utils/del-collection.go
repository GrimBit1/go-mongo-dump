package utils

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DeleteColl(name string) {
	conn, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	db := conn.Database("test")
	db.Collection(name).Drop(context.Background())
}

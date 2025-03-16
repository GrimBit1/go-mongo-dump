package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Find() {
	coll, err := UsersColl()
	if err != nil {
		fmt.Println(err)
		return
	}
	cursor, err := coll.Find(context.Background(), struct{}{}, options.Find().SetSort(map[string]int{
		"first_name": 1,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.Background())

	userSlice := []User{}
	err = cursor.All(context.Background(), &userSlice)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(userSlice) > 0 {
		// Print IP address as a string
		// fmt.Println("IP Address:", userSlice[0].NormalizeIP())
		data, err := json.Marshal(userSlice[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))
	} else {
		fmt.Println("No users found")
	}
}

func FindByID() {
	// If you still want to list all users, call it here.
	ListUsers()

	coll, err := UsersColl()
	if err != nil {

		fmt.Println("This error", err)
		return
	}

	id, err := bson.ObjectIDFromHex("67cdb7acfb3bc8b95b8d7208")
	if err != nil {
		fmt.Println(err)
		return
	}

	var user User
	err = coll.FindOne(context.Background(), bson.M{
		"_id": id,
	}).Decode(&user)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return
	}
	fmt.Printf("%#v\n", user)

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error finding user:", err)
		return
	}
	fmt.Println(string(data))
	// user.NormalizeIP()
	// You may want to do further processing here.
}
func ListUsers() {
	coll, err := UsersColl()
	if err != nil {
		fmt.Println(err)
		return
	}

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		fmt.Println("Error listing users:", err)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			fmt.Println("Decode error:", err)
			continue
		}
		fmt.Printf("%#v\n", user.Id.String())
		return
	}
	if err := cursor.Err(); err != nil {
		fmt.Println("Cursor error:", err)
	}
}

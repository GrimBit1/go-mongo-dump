package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type User struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	Gender    string `json:"gender" bson:"gender"`
	IpAddress net.IP `json:"ip_address" bson:"ip_address"`
}

const uri = "mongodb://root:example@localhost:27017/"

func main() {
	users := []User{}
	data, err := os.ReadFile("./user.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(data, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(len(users))
	timeNow := time.Now().UnixMilli()
	InsertOne(users)
	// MongoRestore()
	fmt.Println(time.Now().UnixMilli()-timeNow, "ms")
	// DeleteColl("users")
}

func InsertMany(users []User) {
	conn, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	db := conn.Database("test")
	usersColl := db.Collection("users")
	res, err := usersColl.InsertMany(context.TODO(), users, options.InsertMany().SetOrdered(false))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Acknowledged)
}
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

func DeleteColl(name string) {
	conn, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println(err)
		return
	}
	db := conn.Database("test")
	db.Collection(name).Drop(context.Background())
}

func MongoImport() {
	cmd := exec.Command("mongoimport", "--uri="+uri, "-d", "test", "-c", "users", "--authenticationDatabase=admin", "--jsonArray", "./user.json")
	fmt.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	fmt.Println("Output:", string(out))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func MongoRestore() {
	cmd := exec.Command("mongorestore", "--uri="+uri, "-d", "test", "-c", "users", "--authenticationDatabase=admin", "--numInsertionWorkersPerCollection", "4", "./users.bson")
	fmt.Println(cmd.String())
	out, err := cmd.CombinedOutput()
	fmt.Println("Output:", string(out))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

}
func WriteUsersToBSONFile(users []User) {
	file, err := os.Create("users.bson")
	if err != nil {
		fmt.Println("Error creating BSON file:", err)
		return
	}
	defer file.Close()

	for i, user := range users {
		data, err := bson.Marshal(user)
		if err != nil {
			fmt.Printf("Error marshaling user [%d]: %v\n", i, err)
			return
		}
		_, err = file.Write(data)
		if err != nil {
			fmt.Printf("Error writing BSON for user [%d]: %v\n", i, err)
			return
		}
	}
	fmt.Println("BSON file written successfully with individual documents")
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"sync"
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
	// InsertMany(users)
	WriteUsersToBSONFile2(users)
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
	wg := sync.WaitGroup{}
	batch := 1
	for i := 0; i < len(users); i += batch {
		wg.Add(1)
		go func() {
			defer wg.Done()
			j := i + batch
			if j > len(users) {
				j = len(users)
			}
			_, err := usersColl.InsertOne(context.TODO(), users[i], options.InsertOne())
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println(res.Acknowledged)
		}()
	}
	wg.Wait()
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

func WriteUsersToBSONFile2(users []User) {
	file, err := os.Create("users.bson")
	if err != nil {
		fmt.Println("Error creating BSON file:", err)
		return
	}
	defer file.Close()

	const workerCount = 1000
	jobs := make(chan User, 100)
	results := make(chan []byte, 100)
	done := make(chan struct{})

	// Worker pool for marshalling users
	var wgWorkers sync.WaitGroup
	for range workerCount {
		wgWorkers.Add(1)
		go func() {
			defer wgWorkers.Done()
			for user := range jobs {
				data, err := bson.Marshal(user)
				if err != nil {
					fmt.Println("Error marshaling user:", err)
					continue
				}
				results <- data
			}
		}()
	}

	// Goroutine to close results when work is done
	go func() {
		wgWorkers.Wait()
		close(results)
		close(done)
	}()

	// Feed jobs
	go func() {
		for _, user := range users {
			jobs <- user
		}
		close(jobs)
	}()

	// Write results to file
	for data := range results {
		_, err := file.Write(data)
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}

	// Wait for the closing goroutine to signal completion
	<-done
	fmt.Println("BSON file written successfully with individual documents")
}

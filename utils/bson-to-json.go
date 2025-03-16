package utils

import (
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/v2/bson"
)

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

package utils

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InsertMany(users []User) {

	usersColl, err := UsersColl()
	if err != nil {
		fmt.Println(err)
		return
	}
	wg := sync.WaitGroup{}
	batch := 10000
	for i := 0; i < len(users); i += batch {
		wg.Add(1)
		go func() {
			defer wg.Done()
			j := i + batch
			if j > len(users) {
				j = len(users)
			}
			_, err := usersColl.InsertMany(context.TODO(), users[i:j], options.InsertMany())
			if err != nil {
				fmt.Println(err)
				return
			}
			// fmt.Println(res.Acknowledged)
		}()
	}
	wg.Wait()
}

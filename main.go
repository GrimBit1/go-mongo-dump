package main

import (
	"encoding/json"
	"fmt"
	"go-mongo-dump/utils"
	"os"
	"time"
)

func main() {
	users := []utils.User{}
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
	utils.FindByRegex()
	// MongoRestore()
	fmt.Println(time.Now().UnixMilli()-timeNow, "ms")
	// DeleteColl("users")
}

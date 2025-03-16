package utils

import (
	"fmt"
	"os/exec"
)

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

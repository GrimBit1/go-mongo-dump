package utils

import (
	"fmt"
	"os/exec"
)

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

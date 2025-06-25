package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("processing cmds..")

	cmd := exec.Command("echo", "I love golang")

	output, err := cmd.Output()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("output:", string(output))

}

package main

import (
	"fmt"
	"os/exec"
	"strings"
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

	//---------------------------------------------

	cmd = exec.Command("grep", "love")

	cmd.Stdin = strings.NewReader("a mother loves her child\nthanks for all the love and support\nloving\nokie")

	output, err = cmd.Output()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("output of reader:-")
	fmt.Println(string(output))


	fmt.Println("----------------------------------------------")

	cmd = exec.Command("sleep", "3")

	err = cmd.Start()
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	err = cmd.Wait()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("after 3 secs")

}

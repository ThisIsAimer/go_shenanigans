package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
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

	cmd = exec.Command("sleep", "2")

	err = cmd.Start()
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	time.Sleep(time.Second * 2)

	// err = cmd.Process.Kill()

	fmt.Println("after 2 secs")

	err = cmd.Wait()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("after 3 secs")

	fmt.Println("-----------------------------------")

	cmd = exec.Command("printenv","GOPATH")

	output,err = cmd.Output()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("env:", string(output))

	fmt.Println("-----------------------------------")


	pr, pw := io.Pipe()

	cmd = exec.Command("grep", "love")

	cmd.Stdin = pr

	go func(){
		defer pw.Close()
		pw.Write([]byte("a mother loves her child\nthanks for all the love and support\nloving\nokie"))
	}()

	output ,err = cmd.Output()
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("the output of pipe is :-")
	fmt.Println(string(output))

	fmt.Println("-----------------------------------------")
	cmd = exec.Command("ls","-l")

	output ,err = cmd.CombinedOutput()

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("ls:", string(output))

}

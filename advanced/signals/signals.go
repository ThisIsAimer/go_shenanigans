package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){

	processId := os.Getpid()
	fmt.Println("process ID:", processId)

	signl := make(chan os.Signal, 1)

	fmt.Println("hello world")

	// SIGINT is signal interrupt : ctrl + c in terminal
	// SIGTERM needs linux and kill -s SIGTERM ProcessId cmd to run
	signal.Notify(signl, syscall.SIGINT, syscall.SIGTERM)

	go func(){
		sig := <- signl
		switch sig{
		case syscall.SIGINT:
			fmt.Println("called signal interrupt")
		case syscall.SIGTERM:
			fmt.Println("called signal terminate")
		}

		fmt.Println("gracefully exiting!")
		os.Exit(1)
	}()

	fmt.Println("working..")
	for{
		time.Sleep(time.Second)
	}
}

// tasklist -lists all tasks
// taskkill /F /PID <PID> for terminating a process
// powershell-
// Stop-Process -Id <PID> -Force
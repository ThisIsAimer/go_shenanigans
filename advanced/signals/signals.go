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
		fmt.Println("recieved signal:",sig)
		os.Exit(1)
	}()

	fmt.Println("working..")
	for{
		time.Sleep(time.Second)
	}
}
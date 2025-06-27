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
	done := make(chan bool, 1)

	fmt.Println("hello world")

	// SIGINT is signal interrupt : ctrl + c in terminal
	// SIGTERM needs linux and kill -s SIGTERM ProcessId cmd to run
	signal.Notify(signl, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func(){
		<- signl
		fmt.Println("Done working! sending done signal")
		done <- true
	}()


	go func(){

		getDone(done)

		fmt.Println("Waiting for next signal!")

		sig := <- signl

		// when we use a signal in a programe, it will over write its original functionality

		switch sig{
		case syscall.SIGINT:
			fmt.Println("called signal interrupt")
		case syscall.SIGTERM:
			fmt.Println("called signal terminate")
		case syscall.SIGHUP:
			fmt.Println("calling signal hangup")
		// this will not compile on windows:

		// case syscall.SIGUSR1:
		// 	fmt.Println("user defined signal1")
		// continue

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




// kill -9 <PID> to kill the process incase of permanent loop
// kill -s SIGINT <PID> it interrupts the task
// kill -s SIGTERM <PID> it terminates the task
// kill -s SIGSTOP <PID> it temperorily stops the task
// kill -s SIGCONBT <PID> it resumes the task


func getDone(done chan bool){
	for{
			select{
			case done := <- done:
				fmt.Println("done Signal sent:", done)
				return
			default:
				fmt.Println("getting the result...")
				time.Sleep(time.Second)
			}
			
		}
}
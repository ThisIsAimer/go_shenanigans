package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(300 * time.Second)
	fmt.Println("timer started!...")

	time.Sleep(1 * time.Second)
	fmt.Println("slept for one second...")

	var stopped bool = false

	//stops timer manually
	stopped = timer.Stop()

	if stopped {
		fmt.Println("timer was stopped manually!")
	}

	fmt.Println("-------------------------------------------------------")

	fmt.Println("reseted timer for 1 sec...")
	timer.Reset(time.Second)

	//<-timer.C is blocking in nature
	<-timer.C

	fmt.Println("end of timer")

	fmt.Println("-------------------------------------------------------")


	//time.after send process itself its like timer.c
	process := time.After(3*time.Second)

	done := make(chan bool)

	go func(){ 
		longProcess() 
		done <- true
		}()

	select{
	case <- process:
		fmt.Println("process is timed out")

	case <- done:
		fmt.Println("long process done!")
	}

	fmt.Println("-------------------------------------------------------")

	timer1 := time.NewTimer(300 * time.Millisecond)
	timer2 := time.NewTimer(300 * time.Millisecond)

	select{
	case <- timer1.C:
		fmt.Println("timer 1 ended")
	case <- timer2.C:
		fmt.Println("timer 1 ended")
	}

	fmt.Println("end of programe")



}


func longProcess(){
	for i := range 20{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

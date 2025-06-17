package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(3 * time.Second)
	fmt.Println("timer started!...")

	time.Sleep(1 * time.Second)
	fmt.Println("slept for one second...")

	var stopped bool = false

	//stops timer manually
	stopped = timer.Stop()

	//<-timer.C is blocking in nature
	if stopped {
		fmt.Println("timer was stopped manually!")
	} else {
		<-timer.C
	}

	fmt.Println("end of programme")

}

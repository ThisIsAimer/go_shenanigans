package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	// non blocking  send operation
	go func() {
		select {
		case channel <- 1:
			fmt.Println("sent value to channel")
		default:
			fmt.Println("channel is not ready to send")
		}
	}()

	//where reciever present, select will send value
	//fmt.Println(<-channel)

	// non blocking recieve operation

	select {
	case number := <-channel:
		fmt.Println("number recieved:", number)

	default:
		fmt.Println("no number recieved")
	}

	fmt.Println("------------------------------------------------------------")

	data := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case value := <-data:
				fmt.Println("recieved value:", value)

			//false is the 0 value for quit in boolean!
			case <-quit:
				fmt.Println("exiting....")
				return

			// case val := <-quit:
			// 	if val {
			// 		fmt.Println("Got true, exit.")
			// 		return
			// 	} else {
			// 		fmt.Println("Got false, ignore.")
			// 	}
			default:
				fmt.Println("waiting for data....")
				time.Sleep(time.Second / 2)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)

	}

	//quit <- false
	quit <- true
	// it keeps looping as the quit channel stays active forever unless there is a
	// reciever

}

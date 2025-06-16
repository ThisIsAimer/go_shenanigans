package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		// when channel is closed, it will recieve the zero value of the data type
		// so we use ok statement to nulify it if closed
		// ok will only be false after channel is closed and empty
		channel1 <- "I love golang"
		close(channel1)

		channel2 <- "I love sushi"
		close(channel2)
	}()

	time.Sleep(time.Second / 10)

	ending := false

	for true {
		select {
		case msg, ok := <-channel1:
			//ignores if channel is closed
			if !ok {
				channel1 = nil
				break
			}

			fmt.Println("recieved message from channel 1:", msg)

		case msg, ok := <-channel2:
			//ignores if channel is closed
			if !ok {
				channel2 = nil
				break
			}

			fmt.Println("recieved message from channel 2:", msg)

		default:
			fmt.Println("no msg found")
			ending = true
		}

		if ending {
			break
		}

	}

	fmt.Println("-------------------------------------------------------------")

	channel := make(chan int)

	go func() {
		time.Sleep(time.Second/2)
		channel <- 1
		close(channel)
	}()

	select {
	case number, ok := <-channel:

		if !ok {
			fmt.Println("channel is closed")
			return
		}

		fmt.Println("we got number:", number)

	// if this is equal or less to the time taken in the groutine
	// then it will time out
	case <-time.After(time.Second):
		fmt.Println("timed out")
	}


	fmt.Println("end of programme!")

}

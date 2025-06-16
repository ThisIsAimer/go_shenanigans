package main

import (
	"fmt"
	"time"
)

func main(){
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func(){
		channel1 <- "I love golang"
		channel2 <- "I love sushi"
	}()

	time.Sleep(time.Second/10)
	
	ending := false

	for true{
		select{
		case msg := <- channel1:
			fmt.Println("recieved message from channel 1:", msg)
		case msg := <- channel2:
			fmt.Println("recieved message from channel 2:", msg)
		default:
			fmt.Println("no msg found")
			ending = true
		}

		if ending{
			break
		}

	}

	fmt.Println("end of programme")

}
package main

import (
	"fmt"
	"time"
)

func main(){
	channel := make(chan string)//this makes a unbuffered channel

	//only unbuffered channels need goroutines
	go func(){
		time.Sleep(2 * time.Second)
		fmt.Println("2 seconds passed")

		channel <- "Hey ther i am here"

		time.Sleep(3 * time.Second)
		fmt.Println("3 seconds passed")
	}()

	//an unbuffered channel will try to imidiately find a reciever of the value
	//being result in this case
	result := <-channel

	//channels will only pause the programmes until they get a value. after it recieves the value the main
	//programme will continue

	fmt.Println(result)

}
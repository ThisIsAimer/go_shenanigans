package main

import (
	"fmt"
	"time"
)

func main() {

	//making buffered channel
	channel := make(chan int, 2)

	//this channel can only hold 2 values of int
	//3 values will cause error
	channel <- 7
	channel <- 10

	fmt.Println("this is buffered channel")

	go func() {
		time.Sleep(3 * time.Second)
		result := <-channel
		fmt.Println(result)

	}()

	//channel will wait for the channel to get empty and then execute
	channel <- 20

	result := <-channel
	fmt.Println(result)

	result = <-channel
	fmt.Println(result)

	//------------------------------------------------------------------------------------
	fmt.Println("--------------------------------------------------------------")

	myChannel := make(chan int, 2)

	go func() {
		time.Sleep(3 * time.Second)
		//myChannel <- 10

	}()
	
	//channels will wait for go routines to end before throwing error
	fmt.Println("the value of channel is", <-myChannel)

}

//note! each line is executed from right to left : end <- start

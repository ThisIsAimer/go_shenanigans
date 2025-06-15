package main

import (
	"fmt"
	"time"
)

func main(){

	//making buffered channel
	channel := make(chan int,2) 

	//this channel can only hold 2 values of int
	//3 values will cause error
	channel <- 7
	channel <- 10

	fmt.Println("this is buffered channel")

	go func(){
		time.Sleep(3 * time.Second)
		result := <- channel
		fmt.Println(result)

	}()

	//channel will wait for the channel to get empty and then execute
	channel <- 20


	result := <- channel
	fmt.Println(result)

	result = <- channel
	fmt.Println(result)

	fmt.Println("--------------------------------------------------------------")

}
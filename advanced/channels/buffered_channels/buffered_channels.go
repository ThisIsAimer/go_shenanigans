package main

import "fmt"

func main(){

	//making buffered channel
	channel := make(chan int,2) 

	//this channel can only hold 2 values of int
	//3 values will cause error
	channel <- 7
	channel <- 10
	//channel <- 20

	fmt.Println("this is buffered channel")

	result := <- channel
	fmt.Println(result)

	result = <- channel
	fmt.Println(result)

}
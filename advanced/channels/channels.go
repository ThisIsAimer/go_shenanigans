package main

import "fmt"

func main(){

	channel := make(chan string) //making a channel

	myString := "I love golang"

	channel <- myString // deadlick as channel is continiously trying to recieve data
						// it stops programme flow

	reciever := <- channel

	fmt.Println(reciever)

}
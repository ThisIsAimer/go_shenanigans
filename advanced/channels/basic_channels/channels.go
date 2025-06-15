package main

import (
	"fmt"
	"time"
)

func main(){

	channel := make(chan string) //making a channel

	myString := "I love golang"


	go func(){
		channel <- myString 
		channel <- "I am soo in love!"
		channel <- "Keep learning!"
		for _, value := range "abcde"{
			channel <- "Alphabet: " + string(value)
		}
	}()

	time.Sleep(1 * time.Second)

	//reciever is in the main go routine
	reciever := <- channel 
	fmt.Println(reciever)

	go func(){
		//it gets a continious stream of data
		reciever = <- channel 
		fmt.Println(reciever)
	}()

	// we need to give it chance to come back to the main thread
    time.Sleep(time.Second/2)

	//or this same would work
	reciever = <- channel 
	fmt.Println(reciever)

	for value := range 5{
	reciever = <- channel 
	fmt.Println(value+1,reciever)
	}

}
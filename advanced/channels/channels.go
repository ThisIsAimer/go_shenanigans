package main

import "fmt"

func main(){

	channel := make(chan string) //making a channel

	myString := "I love golang"


	go func(){
		channel <- myString 
	}()

	reciever := <- channel //reciever is in the main go routine

	fmt.Println(reciever)

}
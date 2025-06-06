package main

import (
	"fmt"
	"log"
)

func main(){

	log.Println("loading....")
	fmt.Println("Golang is ready to execute!")

	log.SetPrefix("Info: ")
	log.Println("This is an info message")
	log.Println("message")

	//flags
	log.SetFlags(log.Ldate)
	log.Println("this is a log message with just date")

	log.SetFlags(log.Ltime)
	log.Println("this is a log message with just time")

	log.SetFlags(log.LstdFlags)
	log.Println("this is a log message with just normal date and time")


}
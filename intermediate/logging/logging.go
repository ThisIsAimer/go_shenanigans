package main

import (
	"fmt"
	"log"
	"os"
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
	log.Println("this is a log message with just stdFlags")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("this is a log message with just normal date and time with file details")

	warnLogger := log.New(os.Stdout,"Waring: ", log.Ldate | log.Ltime)
	newLogger.Println("my custom log")
	warnLogger.Println("my custom log 2")


}

//available for all functions
var(
	newLogger = log.New(os.Stdout,"Log: ", log.Ldate | log.Ltime | log.Lmicroseconds )
)
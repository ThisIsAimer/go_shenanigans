package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

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

	warnLogger := log.New(os.Stdout, "Waring: ", log.Ldate|log.Ltime)
	newLogger.Println("my custom log")
	warnLogger.Println("my custom log 2")


	// os.O_CREATE will create the file if it doesn’t exist,
	// os.O_WRONLY gives you write permission,
	// os.O_APPEND makes every write go to the end of the file.
	
	file, err := os.OpenFile("intermediate/logging/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln("failed to open log file: ", err)
	}

	defer file.Close()

	appLogger := log.New(file, "Waring: ", log.Ldate|log.Ltime)

	appLogger.Println("App is up and running")

}

// available for all functions
var (
	newLogger = log.New(os.Stdout, "Log: ", log.Ldate|log.Ltime|log.Lmicroseconds)
)

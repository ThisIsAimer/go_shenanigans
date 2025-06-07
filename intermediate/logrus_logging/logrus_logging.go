package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main(){
	// "go mod init logrus_logging"for go.mod file
	// "go get github.com/sirupsen/logrus" for package or dependency

	//go.mod must be in the root directory "cd intermediate/logrus_logging"
	fmt.Println("hello world")


	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("this is an info message")
	log.Warn("this is a warning message")
	log.Error("this is an error message")
	log.Println("this is a log")

	log.WithFields(logrus.Fields{
		"username": "bro",
		"method":"get",
	}).Info("log with fields!")

}
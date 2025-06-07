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
	log.Println("this is a log")

}
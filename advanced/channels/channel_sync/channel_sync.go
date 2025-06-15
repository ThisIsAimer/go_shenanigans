package main

import (
	"fmt"
	"time"
)

func main(){

	done := make(chan struct{})

	go func(){
		fmt.Println("working....")
		time.Sleep(2*time.Second) //simulation
		done <- struct{}{}
	}()

	<- done

	fmt.Println("finished working!")


	fmt.Println("---------------------------------------------------------")

	channel := make(chan int)

	go func(){
		channel <- 200
		fmt.Println("sending value:")//blocking until value is recieved

	}()

	value := <- channel //blocking until value is sent
	fmt.Println(value)

}
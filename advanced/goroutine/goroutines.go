package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(2* time.Second)

	fmt.Println("hello from goroutine!")


}

func main(){
	fmt.Println("hello world")

	//the programme ends faster then 2 secs
	//so this code is never executed..
	//if the programme was longer it would
	//go processes it in  another thread
	go sayHello()

}

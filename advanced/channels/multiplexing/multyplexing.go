package main

import "fmt"

func main(){
	channel1 := make(chan int)
	channel2 := make(chan int)

	select{
	case msg := <- channel1:
		fmt.Println("recieved message from channel 1:", msg)
	case msg := <- channel2:
		fmt.Println("recieved message from channel 2:", msg)
	default:
		fmt.Println("no msg recieved!")
	}

}
package main

import "fmt"

func main(){

	channel := make(chan int)

	// non blocking  send operation

	select{
		case channel <- 1:
			fmt.Println("sent value to channel")
		default:
			fmt.Println("channel is not ready to send")
	}

	// non blocking recieve operation

	select{
		case number := <- channel:
			fmt.Println("number recieved:", number)

		default:
			fmt.Println("no number recieved")
	}

	fmt.Println("------------------------------------------------------------")





}
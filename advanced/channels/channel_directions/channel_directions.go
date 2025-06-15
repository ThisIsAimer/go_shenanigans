package main

import (
	"fmt"
)

func main() {

	data := make(chan int)
	//data := make(<- chan int) recieve only channel / only gives out data
	//data := make(chan <- int) send only channel / only takes in data

	//send only goroutine channel
	go func(ch chan<- int) {
		for i := range 5 {
			ch <- i + 1
		}
		close(ch)
	}(data)

	// for value:= range data{
	// 	fmt.Println("Recieved:",value)
	// }

	//or

	// func (ch <- chan int){
	// 	for value:= range ch{
	// 		fmt.Println("Recieved:",value)
	// 	}
	// }(data)

	recieveData(data)

}

//recieve only channel
func recieveData(ch <-chan int) {
	for value := range ch {
		fmt.Println("Recieved:", value)
	}
}

// channels are intermediatories,
// types of channels
// producers(send only)
// consumers(recieve only)

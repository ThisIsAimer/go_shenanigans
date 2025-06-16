package main

import (
	"fmt"
)

func main() {
	//closing channels anly ends sending values to channel
	//closed buffered channels may have values in them that can be accessed!

	channel := make(chan int)

	go func() {
		var i = 10
		for range 5 {
			channel <- i
			i += 10
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Println("recieved value:", value)
	}

	fmt.Println("------------------------------------------------------")

	mydata := make(chan int)
	close(mydata)

	value, ok := <-mydata

	if !ok {
		fmt.Println("channel is closed")
	} else {
		fmt.Println(value)
	}

	fmt.Println("------------------------------------------------------")

	//closing already close channel will cause error

	// ch := make(chan int)

	// go func() {
	// 	close(ch)
	// 	close(ch)
	// }()

	//we should never close channels in 2 different goroutines

	fmt.Println("------------------------------------------------------")

	in := make(chan int)
	out := make(chan int)

	go producer(in)
	go filter(in, out)

	for value := range out{
		fmt.Println("we got value from filter:", value)
	}

}


func producer(channel chan<- int){
	for i := range 5{
		channel <- i*i
	}
	close(channel)
}


func filter(in <-chan int, out chan <-int){
	for value := range in{
		if value % 2 == 0{
			out <- value
		}
	}
	close(out)
}
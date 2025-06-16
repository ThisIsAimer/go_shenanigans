package main

import "fmt"

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

}

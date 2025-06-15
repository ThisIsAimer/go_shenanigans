package main

import (
	"fmt"
	"time"
)

func main() {
	
	done := make(chan struct{})

	go func() {
		fmt.Println("working....")
		time.Sleep(2 * time.Second) //simulation
		done <- struct{}{}
	}()

	<-done

	fmt.Println("finished working!")

	fmt.Println("---------------------------------------------------------")

	channel := make(chan int)

	go func() {
		channel <- 200
		fmt.Println("sending value:") //blocking until value is recieved
	}()

	value := <-channel //blocking until value is sent
	fmt.Println(value)

	fmt.Println("---------------------------------------------------------")

	//---------------------- CHANNEL SYNC IS:- ---------------------------------
	// if we have many channels, we must have equal number of recievers
	// ensuring all the routines are complete

	tasks := make(chan int, 3)

	for id := range 3 {

		//time.Sleep(1 * time.Second)

		go func(i int) {
			fmt.Println("Goroutine working on:", id+1, "task")
			tasks <- id + 1 //sending signal of completion
		}(id)

	}

	for range 3 {
		result := <-tasks // waiting for all goroutines signal completion
		fmt.Println(result)
	}
	// 3 or 2 or 1 may execute first with no order cause of go runetime threading
	//as they executing really fast
	//with time.Sleep they will execute on order

	fmt.Println("all goroutines are finished")

	fmt.Println("---------------------------------------------------------")
	//continious stream of data

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- fmt.Sprintf("we got \"%d\" no. of data", i)
			time.Sleep(100 * time.Microsecond)
		}
		//will close the channel telling that the channel will not recieve more data
		close(data)
	}()

	// ranging over channel!
	for value := range data {
		fmt.Println("we got:", value, "at", time.Now())
	}
	// loops over the channel until coming across close function

}

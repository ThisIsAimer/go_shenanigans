package main

import (
	"fmt"
	"time"
)

type ticketRequest struct {
	personId   int
	numTickets int
	cost       int
}

func main() {
	WorkerNum := 3
	numTasks := 10

	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	// made go routines as num of workers
	for employee := range WorkerNum {
		go worker(employee+1, tasks, results)
	}

	//giving tasks
	for i := range numTasks {
		tasks <- i
	}

	close(tasks)

	for range numTasks {
		result := <-results

		fmt.Println("result", result)
	}

	fmt.Println("-----------------------------------------------")

	numRequests := 5
	cost := 10

	ticketReq := make(chan ticketRequest, numRequests)
	processedReq := make(chan int, numRequests)

	for range 3 {
		go ticketProcessor(ticketReq, processedReq)
	}

	for i := range numRequests {
		ticketReq <- ticketRequest{personId: i + 1, numTickets: (i + 1) * 2, cost: ((i + 1) * 2) * cost}
	}
	close(ticketReq)

	for range numRequests {
		processed := <-processedReq

		fmt.Println("processing done for person id:", processed)
	}
	close(processedReq)

}

func worker(id int, tasks <-chan int, results chan<- int) {

	// range tasks can dynamically increase or decrease
	for task := range tasks {
		fmt.Printf("worker %d is working on task %d...\n", id, task)

		time.Sleep(time.Second / 4)

		results <- task*2 + 1
	}

}

func ticketProcessor(requests <-chan ticketRequest, result chan<- int) {
	for request := range requests {
		fmt.Printf("processing %d tickets of personId %d, total cost: %d \n", request.numTickets, request.personId, request.cost)
		time.Sleep(time.Second)

		result <- request.personId
	}
}

package main

import (
	"fmt"
	"time"
)

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

}

func worker(id int, tasks <-chan int, results chan<- int) {

	// range tasks can dynamically increase or decrease
	for task := range tasks {
		fmt.Printf("worker %d is working on task %d...\n", id, task)

		time.Sleep(time.Second)

		results <- task*2 + 1
	}

}

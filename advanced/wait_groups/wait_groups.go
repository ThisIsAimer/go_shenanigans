package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	numWorkers := 3

	//waits for i no. of go routines to finish at wg.wait
	wg.Add(numWorkers)

	for i := range numWorkers {
		go worker(i+1, &wg)
	}

	// only continues when wg.Add(i) i no. of go routines are executed!
	wg.Wait()

	fmt.Println("end of programme!")
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d starting \n", id)

	time.Sleep(time.Second)

	fmt.Printf("worker %d has finished the task \n", id)
}

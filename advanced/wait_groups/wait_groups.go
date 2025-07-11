package main

import (
	"fmt"
	"sync"
	"time"
)

type majdur struct {
	id   int
	work string
}

func (w *majdur) doWork(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%d id. worker is doing %s task\n", w.id, w.work)
	time.Sleep(time.Second)

	fmt.Printf("%s task done\n", w.work)
}

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

	fmt.Println("--------------------------------------------------")

	var wGroup sync.WaitGroup
	numEmp := 3
	numTasks := 5

	results := make(chan int, numTasks)
	tasks := make(chan int, numTasks)

	wGroup.Add(numEmp)

	for i := range numEmp {
		go working(i+1, tasks, results, &wGroup)
	}

	for i := range numTasks {
		tasks <- i + 1
	}
	close(tasks)

	// closing in goroutine
	// results will only close when all go routines are done!
	go func() {
		wGroup.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}

	fmt.Println("-----------------------------------------")

	var constructionWait sync.WaitGroup

	constructionWork := []string{"Placing bricks", "painting", "buying furniture", "building walls", "Cleaning up"}

	for i, value := range constructionWork {
		MyMajdur := majdur{id: i + 1, work: value}
		constructionWait.Add(1)

		go MyMajdur.doWork(&constructionWait)
	}

	constructionWait.Wait()

	fmt.Println("all works are complete!")

}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d starting \n", id)

	time.Sleep(time.Second / 4)

	fmt.Printf("worker %d has finished the task \n", id)
}

func working(id int, tasks <-chan int, result chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("worker %d is working on %d no. task..\n", id, task)
		time.Sleep(time.Second / 4)

		result <- task * 2
		fmt.Printf("worker %d is finished task no.%d...\n", id, task)
	}

}

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*

Concurrency
• The ability of a system to handle multiple tasks simultaneously. It involves
 managing multiple tasks that are in progress at the same time but not necessarily
 executed at the same instant.

Parallelism
• The simultaneous execution of multiple tasks, typically using multiple processors
or cores, to improve performance by running operations at the same time.

*/

func printNumbers(number int) {

	for i := range number {

		fmt.Println(time.Now())
		fmt.Println("number is:", i)
		time.Sleep(time.Second / 10)

	}
}

func printLetters(word string) {

	for _, letter := range word {
		fmt.Println(time.Now())
		fmt.Println("letter is:", string(letter))
		time.Sleep(time.Second / 10)

	}

}

func HeavyTask(id int, wg *sync.WaitGroup){
	defer wg.Done()

	fmt.Printf("task %d has started\n", id)

	for range 1_00_00_000{

	}

	fmt.Println(time.Now())

	fmt.Printf("task %d is finished\n", id)
}


func main() {

	fmt.Println("Programme:")

	//these are concurrent, but seem parallel due to my processors speed
	go printNumbers(10)
	go printLetters("golang")

	time.Sleep(time.Second*2)

	fmt.Println("-----------------------------")
	var wg sync.WaitGroup

	numThreads := 6
	runtime.GOMAXPROCS(numThreads)


	for i := range numThreads{
		wg.Add(1)
		go HeavyTask(i+1, &wg)
	}

	wg.Wait()

}

package main

import (
	"fmt"
	"sync"
)

type counter struct{
	mutex sync.Mutex
	count int
}

func (c *counter) incriment(){
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.count++
}

func (c *counter) getValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.count
}



func main(){
	// mutexes disallow other go routines to temper with already running functions

	var wg sync.WaitGroup

	myCounter := &counter{}

	numGoroutines := 10


	for range numGoroutines{
		wg.Add(1)
		go func(){
			defer wg.Done()
			for range 1000{
				myCounter.incriment()
			}
		}()
	}

	wg.Wait()

	fmt.Println("the final counter is", myCounter.getValue())

}
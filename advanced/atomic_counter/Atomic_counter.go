package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// atomic works like mutexes. it makes sure that shared variables in goroutines
// are updated properly! we are free from locking and unlocking manually

type atomicCounter struct{
	count int64
}

func (ac *atomicCounter) incriment(){
	atomic.AddInt64(&ac.count, 1)
}

func (ac atomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func main(){
	var waitgroup  sync.WaitGroup

	numGoroutines := 10

	counter := &atomicCounter{}

	for range numGoroutines{
		waitgroup.Add(1)
		go func ()  {
			defer waitgroup.Done()

			for range 1000{
				counter.incriment()
			}
		}()
	}

	waitgroup.Wait()

	fmt.Println("The counter is now:", counter.getValue())

}
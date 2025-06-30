package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMux sync.RWMutex
var counter int

func readCounter (wg *sync.WaitGroup){
 defer wg.Done()
 rwMux.RLock()
 fmt.Println("counter is:", counter)
 rwMux.RUnlock()
}

func writeCounter (wg *sync.WaitGroup, value int){
 defer wg.Done()
 rwMux.Lock()
 counter = value
 fmt.Printf("written value %d for counter\n", counter)
 rwMux.Unlock()

}


func main() {

	var wg sync.WaitGroup

	for range 5{
		wg.Add(1)
		go readCounter(&wg)
	}



	wg.Add(1)
	// time.Sleep(time.Second)
	writeCounter(&wg, 77)


	time.Sleep(time.Millisecond)


	

}

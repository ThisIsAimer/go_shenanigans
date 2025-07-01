package main

import (
	"fmt"
	"sync"
	"time"
)

const bufferSize = 5

type buffer struct {
	items []int
	mux   sync.Mutex
	cond  *sync.Cond
}

func newBuffer(size int) *buffer {
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mux)

	return b
}

func (b *buffer) produce(item int) {
	b.mux.Lock()
	defer b.mux.Unlock()

	// we use for loop cause wait can wait up due to other conditions like broadcast
	// and we dont want it to execulte without context
	
	for len(b.items) >= bufferSize {
		// what it does is it unloacks and locks the function continiously so that
		// other go routines can access the elements and modify them
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("produced a  new item:", item)
	b.cond.Signal()

}

func (b *buffer) consume() int {
	b.mux.Lock()
	defer b.mux.Unlock()

	for len(b.items) == 0 {
		// if 0 it stops and waits for other func to append to slice
		b.cond.Wait()
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("consumed:", item)
	b.cond.Signal()
	return item
}

func producer(buf *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		buf.produce(i + 100)
		time.Sleep(time.Millisecond * 100)
	}
}

func consumer(buf *buffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 10 {
		buf.consume()
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {
	buffer := newBuffer(bufferSize)
	var wg sync.WaitGroup

	wg.Add(2)
	go producer(buffer, &wg)
	go consumer(buffer, &wg)

	wg.Wait()
}

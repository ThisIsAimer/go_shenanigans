package main

import (
	"fmt"
	"sync"
)


const bufferSize = 5

type buffer struct{
	items []int
	mux sync.Mutex
	cond *sync.Cond
}


func newBuffer(size int) *buffer{
	b := &buffer{
		items: make([]int, 0, size),
	}
	b.cond = sync.NewCond(&b.mux)

	return b
}

func (b *buffer) produce(item int){
	b.mux.Lock()
	defer b.mux.Unlock()

	for len(b.items) >= bufferSize{
		b.cond.Wait()
	}

	b.items = append(b.items, item)
	fmt.Println("produced a new item:", item)
	b.cond.Signal()
	
}

func (b *buffer) consume() int{
	b.mux.Lock()
	defer b.mux.Unlock()

	for len(b.items)  == 0 {
		// if 0 it stops and waits for other func to append to slice
		b.cond.Wait()
	}
	item := b.items[0]
	b.items = b.items[1:]
	fmt.Println("consumed:",item)
	b.cond.Signal()
	return item
}

func main(){
	fmt.Println("new programe")

}
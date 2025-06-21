package main

import (
	"fmt"
	"time"
)

type statefulWorker struct{
	count int
	ch chan int
}

func (w *statefulWorker) start(){
	go func(){
		for {
			select{
			case value := <- w.ch:
				w.count+= value
				fmt.Println("current count:", w.count)
			default:
				fmt.Print("")
			}

		}

	}()
}

func (w *statefulWorker) send (value int){
	w.ch <- value
}


func main(){
	stWorker := statefulWorker{ch: make(chan int)}

	stWorker.start()

	for range 20{
		stWorker.send(20)
		time.Sleep(time.Second)
	}

}
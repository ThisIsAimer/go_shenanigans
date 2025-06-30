package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var mux1, mux2 sync.Mutex

	// deadlock is created as goroutine 2 tries to lock mux1 when its already locked
	go func () {
		mux1.Lock()
		fmt.Println("goroutine 1 mutex 1 locked")
		time.Sleep(time.Second)
		mux2.Lock()
		fmt.Println("goroutine 1 mutex 2 locked")
		mux1.Unlock()
		mux2.Unlock()
	} ()

	// deadlock is created as goroutine 1 tries to lock mux2 when its already locked
	go func () {
		mux2.Lock()
		fmt.Println("goroutine 2 mutex 2 locked")
		time.Sleep(time.Second)
		mux1.Lock()
		fmt.Println("goroutine 2 mutex 1 locked")
		mux2.Unlock()
		mux1.Unlock()
	} ()

	time.Sleep( time.Second *3 )
	fmt.Println("programme completed")

	/*
	//this doesnt create a deadlock
	go func () {
		mux1.Lock()
		fmt.Println("goroutine 1 mutex 1 locked")
		time.Sleep(time.Second)
		mux2.Lock()
		fmt.Println("goroutine 1 mutex 2 locked")
		mux1.Unlock()
		mux2.Unlock()
	} ()

	go func () {
		mux1.Lock()
		fmt.Println("goroutine 1 mutex 1 locked")
		time.Sleep(time.Second)
		mux2.Lock()
		fmt.Println("goroutine 1 mutex 2 locked")
		mux1.Unlock()
		mux2.Unlock()
	} ()
	*/

}
package main

import (
	"fmt"
	"sync"
)

var once sync.Once


func initialize(){
	fmt.Println("this will execute only once!")
}

func main(){
	var wg sync.WaitGroup

	for i := range 5{
		wg.Add(1)
		go func(){
			defer wg.Done()
			fmt.Println("goroutine no.",i+1)
			// will only execute once
			once.Do(initialize)
		}()
	}

	wg.Wait()
	fmt.Println("end of function")

}
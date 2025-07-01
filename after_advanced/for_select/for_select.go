package main

import (
	"fmt"
	"time"
)

func main(){

	fmt.Println("Ticking!")

	ticker := time.NewTicker(time.Second)

	quit := make(chan string)

	go func(){
		time.Sleep(time.Second * 5)
		close(quit)
	}()

	for{
		select{
		case <-ticker.C:
			fmt.Println("tick")
		case <-quit:
			fmt.Println("ending program")
			return
		}
	}

}
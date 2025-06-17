package main

import (
	"fmt"
	"time"
)

func main(){
	ticker := time.NewTicker(time.Second/2)

	var num int = 0


	for {
		num++
		fmt.Println(<- ticker.C)
		fmt.Println("currently ticking:", num)
		if num == 4{
			ticker.Stop()
			break
		}
	}

	fmt.Println("-----------------------------------------------------")

	myTicker := time.NewTicker(time.Second)
	stop := time.After(time.Second*5)


	processing(myTicker,stop)

}


func processing(ticker *time.Ticker, timer <-chan time.Time){
	for{
		select{
		case tick := <-ticker.C:
			fmt.Println("processing:", tick)
		case <-timer:
			fmt.Println("timeout!")
			return
		}
	}
}
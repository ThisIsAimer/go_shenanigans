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
		if num == 10{
			ticker.Stop()
			break
		}
	}

}
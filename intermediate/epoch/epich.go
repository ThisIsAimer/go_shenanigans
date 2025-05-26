package main

import (
	"fmt"
	"time"
)

func main(){
	//epoch 00:00:00 UTC  jan1 1970
	var now = time.Now()
	fmt.Println("time now:",now)
	var unixTime = now.Unix()
	fmt.Println("unix time now:",unixTime)

	//second argument is nano-seconds
	t := time.Unix(unixTime,0)

	fmt.Println("time:" , t)
	fmt.Println("time formatted:" , t.Format("02-04-2006"))
}
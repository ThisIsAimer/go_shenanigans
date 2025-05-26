package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())

	istLocation, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		panic(err)
	}

	specificTime := time.Date(2025,time.February,06,0,0,0,0,istLocation)
	//we could use time.UTC

	fmt.Println("specific time:", specificTime)
}
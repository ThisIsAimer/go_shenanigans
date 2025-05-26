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

	//the initial time must be Mon jan 2 15:04:05 mst 2006
	//turns string to time
	var parsedTime,_ = time.Parse("2006-01-02","2020-06-09")
	fmt.Println(parsedTime)

	t := time.Now()
	fmt.Println("formatted time:", t.Format("Mon 02-01-06 15:04"))
	var addedTime = t.Add(time.Hour * 24)
	fmt.Println("one day later:", addedTime)
	fmt.Println("formatted time:", addedTime.Format("Mon 02-01-06 15:04"))
	fmt.Println("formatted time:", addedTime.Weekday())
}
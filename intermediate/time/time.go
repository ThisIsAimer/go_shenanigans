package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())

	//custom time
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
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("formatted time:", t.Format("Mon 02-01-06 15:04"))
	var addedTime = t.Add(time.Hour * 24)

	fmt.Println("one day later:", addedTime)
	fmt.Println("formatted time:", addedTime.Format("Mon 02-01-06 15:04"))
	fmt.Println("formatted time:", addedTime.Weekday())


	fmt.Println("rounded time:", addedTime.Round(time.Hour))

	//converting timezone into other
	istLocation, err = time.LoadLocation("Asia/Kolkata")

	if err != nil {
		panic(err)
	}

	var t1 = time.Date(2025,time.February,06,0,0,0,0,time.UTC)

	var tLocal = t1.In(istLocation)
	fmt.Println("-----------------------------------------------------------------")
	fmt.Println("time (UTC) converted to local:",tLocal)

	//it rounds off down time and will not go to next hour
	fmt.Println("truncated time:",time.Now().Truncate(time.Hour))


	//time in new york
	yorkTime, err := time.LoadLocation("America/New_York")
	if err != nil {
		panic(err)
	}
	fmt.Println("time in new york:", time.Now().In(yorkTime))

	fmt.Println("-------------------------------------------------------------")
	//subtracted time
	time1 := time.Date(2025, time.May, 26, 10, 0, 0, 0, time.UTC)
	time2 := time.Date(2025, time.May, 26, 9, 30, 0, 0, time.UTC)

	// Subtract t2 from t1
	diff := time1.Sub(time2)

	fmt.Println("sub time difference:", diff) 
	fmt.Println("time 1 after time 2:", time1.After(time2)) 
	

}
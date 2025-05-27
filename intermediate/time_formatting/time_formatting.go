package main

import (
	"fmt"
	"time"
)

func main() {
	// Mon Jan 02 15:04:05 MST 2006
	layout := "2006-01-02T15:04:05Z07:00"
	myString := "2025-05-27T12:02:06Z"
	var t, err = time.Parse(layout, myString)

	if err != nil {
		fmt.Println("error parsing time", err)
		return
	}
	fmt.Println(t)

	//custom layout according to data
	var myString1 = "Tue May 05 2025 03:02 PM"
	var layout1 = "Mon Jan 02 2006 03:04 PM"

	var t1, err1 = time.Parse(layout1, myString1)

	if err1 != nil {
		fmt.Println("error parsing time", err)
		return
	}
	fmt.Println(t1)

}

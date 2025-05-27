package main

import (
	"fmt"
	"strconv"
)

func main(){
	var a = "1234"
	intA,err := strconv.Atoi(a)
	if err != nil{
		panic(err)
	}

	fmt.Println("converted string:", intA+10)

}
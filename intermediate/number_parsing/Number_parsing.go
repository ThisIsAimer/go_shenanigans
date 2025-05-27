package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = "1234"
	intA, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}

	fmt.Println("converted string:", intA+10)

	//specifit size

	parseInt, err := strconv.ParseInt(a, 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Println("parsed int:", parseInt)


	var pi = "3.14"
	decPi , err := strconv.ParseFloat(pi,32)
	if err != nil {
		panic(err)
	}

	fmt.Printf("value of pi: %0.2f\n",decPi)


	//binarystring to int
	binaryStr := "10101"
	parseBi,err := strconv.ParseInt(binaryStr,2,32)
	if err != nil {
		panic(err)
	}
	fmt.Println("value of binary is", parseBi)
	fmt.Println("value of binary is", fmt.Sprint(parseBi))

	//hexadecimal
	hexStr := "ff"
	parsehex,err := strconv.ParseInt(hexStr,16,32)
	if err != nil {
		panic(err)
	}
	fmt.Println("value of Hexadec is", parsehex)


	invalidStr := "123abc"
	parseInvalid,err := strconv.Atoi(invalidStr)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	fmt.Println("value of invalid is", parseInvalid)


}

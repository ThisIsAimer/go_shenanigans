package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	var str1 = "hello"
	var str2 = "world"

	fmt.Println(str1 +" "+ str2)
	fmt.Println(str1[1])//ascii
	result := str1+" "+ str2

	fmt.Println(result[1:7])
	//-----------------------------------------------------------------------
	num := 77
	var numToString = strconv.Itoa(num)
	fmt.Println(reflect.TypeOf(numToString))

	fruits := "Apple, Banana, Orange"
	basket := strings.Split(fruits,", ")
	fmt.Println(basket)

	var countries = []string{"India","japan","russia"}
	var myCounties = strings.Join(countries,", ")
	fmt.Println(myCounties)
	fmt.Println(strings.Contains(myCounties, "India"))

	greetings := strings.Replace(result,"world","India",1)
	println(greetings)

	// it trims "   Hello world!  "
	// to "hello world!"
	fmt.Println(strings.TrimSpace(greetings)) 

	fmt.Println(strings.ToLower(greetings))
	fmt.Println(strings.ToUpper(greetings))

	fmt.Println(strings.Repeat("batman ", 4))
	fmt.Println(strings.Count(greetings,"India"))
	fmt.Println(strings.HasPrefix(greetings,"he"))

}

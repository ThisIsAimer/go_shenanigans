package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	var str1 = "hello"
	var str2 = "world"

	fmt.Println(str1 +" "+ str2)
	fmt.Println(str1[1])//ascii
	result1 := str1+" "+ str2

	fmt.Println(result1[1:7])
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

	greetings := strings.Replace(result1,"world","India",1)
	println(greetings)

	// it trims "   Hello world!  "
	// to "hello world!"
	fmt.Println(strings.TrimSpace(greetings)) 

	fmt.Println(strings.ToLower(greetings))
	fmt.Println(strings.ToUpper(greetings))

	fmt.Println(strings.Repeat("batman ", 4))
	fmt.Println(strings.Count(greetings,"India"))
	fmt.Println(strings.HasPrefix(greetings,"he"))

	string1 := "hello 123 go! 12.12 32 54"
	regular := regexp.MustCompile(`\d+`) //finds ints in the para
	matches := regular.FindAllString(string1,-1)// -1 finds all
	fmt.Println(matches)

	//------------------------------------------------------------------
	// STRING BUILDER is used to build strings in a memory efficient way
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("world!")

	// Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh!")
	result = builder.String()
	fmt.Println(result)
}

package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("hello world")
	//regular expressions are used to find things in a string

	//\. for .
	// - needs to be at the end
	//we can add []\+ for actual + symbol
	var re = regexp.MustCompile(`[a-zA-Z1-9._-]+@[a-zA-Z1-9.]+\.[a-zA-Z]{2,}`)

	var email1 = "example@email.com"
	var email2 = "invalid_email"

	fmt.Println("email1: ", re.MatchString(email1))
	fmt.Println("email1: ", re.MatchString(email2))

	re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	date := "2032-02-06"

	var subMatch = re.FindStringSubmatch(date)
	fmt.Println(subMatch)
	fmt.Println(subMatch[0])
	fmt.Println(subMatch[1])
	fmt.Println(subMatch[2])
	fmt.Println(subMatch[3])

	//must be in [] brackets when taking a bunch of letters
	re = regexp.MustCompile(`[aeiou]`)
	//re = regexp.MustCompile(`e`)
	var hello = "hello world"
	var replaced = re.ReplaceAllString(hello, "*")

	fmt.Println(replaced)

	//flags in regular expressions

	// i- case insensitive (caps or small)
	// m- multi line model
	// s- dot matches

	// flags must initialise with ?
	re = regexp.MustCompile(`(?i)go`)

	myString := "Golang is amazing"
	// its accepting capital g (G)
	println("match:", re.MatchString(myString))

}

package main

import (
	"fmt"
	"math/rand"
)

func main(){

	//rand.intn generates a random number by itself from 0 to the given number
	//autoseeding
	var randomNumber = rand.Intn(101)

	fmt.Println(randomNumber)

}
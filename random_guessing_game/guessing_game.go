package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var source = rand.NewSource(time.Now().UnixNano())//creates source/seed
	var random = rand.New(source)//creates a random number gen using the seed

	//generating a number between 1-100
	var randomNumber = random.Intn(100) + 1

	fmt.Println("welcome to the number guessing game!")
	fmt.Println("now you must enter a number to see if you are correct!")

	var guess int

	for {
		println("-------------------------------------------------------")
		fmt.Print("guess number: ")
		fmt.Scanln(&guess)

		if guess > randomNumber {
			println("Too high!, guess lower")
		} else if guess < randomNumber {
			println("Too low!, guess higher ")
		} else {
			println("your guessed was correct!")
			break
		}
	}

}

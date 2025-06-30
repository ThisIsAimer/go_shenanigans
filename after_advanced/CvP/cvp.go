package main

import (
	"fmt"
	"time"
)

/*

Concurrency
• The ability of a system to handle multiple tasks simultaneously. It involves
 managing multiple tasks that are in progress at the same time but not necessarily
 executed at the same instant.

Parallelism
• The simultaneous execution of multiple tasks, typically using multiple processors
or cores, to improve performance by running operations at the same time.

*/

func printNumbers(number int) {

	for i := range number {

		fmt.Println(time.Now())
		fmt.Println("number is:", i)
		time.Sleep(time.Second / 2)

	}
}

func printLetters(word string) {

	for _, letter := range word {
		fmt.Println(time.Now())
		fmt.Println("letter is:", string(letter))
		time.Sleep(time.Second / 2)

	}

}

func main() {

	fmt.Println("Programme:")

	//these are concurrent, but seem parallel due to my processors speed
	go printNumbers(10)
	go printLetters("golang")

	time.Sleep(time.Second * 5)

}

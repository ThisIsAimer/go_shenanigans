package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func sayHello() {
	time.Sleep(2* time.Second)

	fmt.Println("hello from goroutine!")


}

func printNumbers(n int){

	var i = 0

	for i<n {

		fmt.Println(i)

		time.Sleep(2 * time.Second)

		i++
	}
}


func printLetters(str string){
	for _,element := range str{
		fmt.Println(string(element))
		time.Sleep(2 * time.Second)
		fmt.Println("---------------------------------")
	}
}


func doSmth() error{
	time.Sleep(1* time.Second)

	return fmt.Errorf("an error occured")
}


func main(){
	fmt.Println("hello world")

	var err error

	//the programme ends faster then 2 secs
	//so this code is never executed..
	//if the programme was longer it would
	//go processes it in  another thread

	go func(){
		err = doSmth()
	}()
	// err = go doSmth() this doesnt work


	go sayHello()
	go printNumbers(10)
	go printLetters("ABCDEFGHIJKLMNOP")

	//we have to wait more then 1 sec for the error to occur from go routine or the error will not be registered
	time.Sleep(2*time.Second)

	if err != nil {
		fmt.Println("error is:", err)
	}// for do smth

	fmt.Println("enter any sentance: ")


	// fmt.scanln() only reads one word till space
	// we use bufio to read a line

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	
	if err != nil {
		fmt.Println("error is:", err)
		return
	}


	fmt.Println("your variable is:", input)

}

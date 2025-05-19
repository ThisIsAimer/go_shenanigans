package main

import "fmt"

func main() {
	process()

	fmt.Println("process end")

}

//if panic recover will recover the programme
//so it doesnt crash

//all defer will take process weather before or after the recover statement
func process() {
	defer fmt.Println("first defer")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover:", r)
		}
	}()
	defer fmt.Println("second defer")

	panic("something went wrong")
}

package main

import "fmt"

func main() {
	var array [3]int = [3]int{2, 3, 4}
	var copied_array *[3]int = &array
	copied_array[1] = 10

	fmt.Println(*copied_array)//as copied_array is a pointer array we need to dereference it

}
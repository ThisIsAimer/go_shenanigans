package main

import (
	"fmt"
	"slices"
)

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	var slice = a[1:6]
	// a[:]makes a slice

	slice = append(slice, a[2], 10)
	slice = append(slice, a[:]...)

	//a[:]... unpacks the array into indidual elements

	var slice_copy = make([]int, len(a))
	copy(slice_copy, a[:])

	//or

	// var slice_copy [6]int
	// copy(slice_copy[:],a[:])

	fmt.Println(slice)
	fmt.Println(slice_copy)

	var slice_equal = slices.Equal(slice, slice_copy)

	fmt.Println("slice and slice copy equal:", slice_equal)

	var twoD = make([][]int, 4)

	for i := 0; i < 4; i++ {
		var innerLength = i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	slice = slice[2:4]

	fmt.Println("capacity of slice is", cap(slice))
	// slice is a portion of array amd cap tells the capacity of the array

	fmt.Println("length of slice is", len(slice))
	//length is how big the array is!

}

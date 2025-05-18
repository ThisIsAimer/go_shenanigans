package main
import "fmt"

func main(){

	//var numbers = []int{1,2,3,4,5,6,7}
	var numbers = [8]int{1,2,3,4,5,6,7,8}

	//fmt.Println("sum of nums are", sum(numbers...))
	// fmt.Println("sum of nums are", sum(numbers[:]...))
	// fmt.Println("sum of 5,6,7 are", sum(5,6,7))
	 fmt.Println("sum of nums are", sum(numbers))

}

//variadic parameters must be the last parameter of a function
// 

func sum(nums [8]int) int {
	var i int = 0

	for _, v := range nums {
		i += v
	}

	return i
}
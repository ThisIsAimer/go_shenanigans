package main
import "fmt"

func main(){

	//var numbers = []int{1,2,3,4,5,6,7}
	var numbers = [8]int{1,2,3,4,5,6,7,8}

	//fmt.Println("sum of nums are", sum(numbers...))
	// fmt.Println("sum of nums are", sum(numbers[:]...))
	// fmt.Println("sum of 5,6,7 are", sum(5,6,7))
	s, n := sum(numbers,"raunak")
	 fmt.Printf("sum of nums are %d, %s", s, n )

}

//variadic parameters must be the last parameter of a function
// 

// func sum(nums ...int) (int){
// 	var i int = 0
//
// 	for _, v := range nums{
// 		i += v
// 	}
//
// 	return i
// }

func sum(nums [8]int, name string) (int ,string) {
	var i int = 0

	for _, v := range nums {
		i += v
	}

	return i, name
}
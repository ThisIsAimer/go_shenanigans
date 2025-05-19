package main

import "fmt"

func main() {
	fmt.Println("factorial is:", factorial(5))
	fmt.Println("factorial is:", factorial(10))

	fmt.Println("factorial loop is:", factorialLoop(5))

	fmt.Println("sumof figits is:", sumOfDigits(123))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func sumOfDigits(num int) int{
	if num < 10{
		return num
	}else{
		return num%10 + sumOfDigits(num / 10)
	}
}

// in loop
func factorialLoop(n int) int {
	var i int = 1

	for n > 0 {
		i *= n
		n--
	}

	return i
}

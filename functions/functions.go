package main
import "fmt"

func main(){
	var add = addition
	var value = add(1,3)
	fmt.Println("addition =",value)


	value = operation(9,7,addition)
	fmt.Println("operation =",value)

	var multipliableValue = multiplicationBy(7)
	fmt.Println(multipliableValue(7))
	fmt.Println(multiplicationBy(7)(10))

}

func addition(a,b int) int {
	return a+b
}

func operation(a,b int, operation func(c,d int) int) int {
	return operation(a,b)
}

// go most likely has a built in storage of the passed parameters
func multiplicationBy(a int) func(b int) int{
	return func(b int) int{
		return a*b
	}
}
package main
import "fmt"

func main(){
	
	// we cant use i as var in for loop
	// we need to use the gofer syntax
	for i := 0; i<=10; i++ {
		fmt.Println(i)
	}

	var numbers []int = []int{9, 3, 3, 4, 8, 6}
	//numbers := []int {1,2,3,4,5}
	println("---------------------------------------------------")
	//  index, value
	for _, value := range numbers{
		println(value)
	}


	println("---------------------------------------------------")

	for value := range 20 {

		if value %2 == 1 {
			continue
		}

		println(value)

		if value == 10 {
			break
		}

	}

	rows := 5;
	println("---------------------------------------------------")

	for i := 1; i<=rows ; i++{
		for j := 1; j<=rows-i;j++ {
			fmt.Print(" ")
		}
		for k := 1; k<=2*i-1;k++{
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
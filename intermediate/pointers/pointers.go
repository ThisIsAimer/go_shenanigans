package main
import "fmt"

func main(){
	var a = 10
	//initialise a pointer
	var ptr *int = &a

	fmt.Println("value of a", *ptr)
	ptrIncriment(ptr)
	fmt.Println("value of a", a)

}

func ptrIncriment(ptr *int) {
	*ptr++
}
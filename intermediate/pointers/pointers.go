package main
import "fmt"

func main(){
	var a = 10
	//initialise a pointer
	var ptr *int = &a

	fmt.Println("value of a", *ptr)
	ptr_incriment(ptr)
	fmt.Println("value of a", a)

}

func ptr_incriment(ptr *int) {
	*ptr++
}
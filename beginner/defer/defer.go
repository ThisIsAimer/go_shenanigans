package main
import "fmt"

func main(){
	process(10)

}

func process(i int) {
	//first in last out
	defer fmt.Println("first value of i is:", i)
	defer fmt.Println("second")
	defer fmt.Println("third")
	i= 100
	fmt.Println("last value of non defer i is:", i)
	fmt.Println("-----------------------------------------------------")
	defer fmt.Println("last value of i is:", i)
}
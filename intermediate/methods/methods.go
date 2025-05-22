package main
import "fmt"

func main(){
	fmt.Println("hello world")

	var i = myInt(1)
	fmt.Println(i)
	fmt.Println(i.addition(2))
	fmt.Println(i.welcome())

}

type myInt int

func (i myInt) addition(n int) int {
	return int(i) + n
}

func (myInt) welcome() string {
	return "whats up"
}
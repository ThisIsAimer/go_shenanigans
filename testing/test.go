package main
import "fmt"


func main(){
	var stack = []any{
		"hey", 23, 9.3,'c',
	}
	fmt.Println(stack...)
}
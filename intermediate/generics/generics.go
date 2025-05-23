package main
import "fmt"

type stack[T any] struct{
	elements []T
}

func (s *stack[T]) push(element T) {
	s.elements = append(s.elements, element)
}

func (s *stack[T]) pop() (T, bool){

	if len(s.elements) == 0{
		var zero T
		return zero, false
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]//deletes last element
	return element , true
}

func (s stack[T]) isEmpty() bool {
	return len(s.elements) == 0
}
func(s stack[T]) printElements() {
	if len(s.elements) == 0{
		fmt.Println("the stack is empty")
	}else{
		for _, element := range s.elements{
			fmt.Print(element , " ")
		}
		println()
	}
}


func main(){

	//shouldnt name struct's instance name same as struct's name
	var intStack = stack[int]{}
	intStack.push(20)
	intStack.push(50)
	intStack.push(100)

	intStack.printElements()
	fmt.Println(intStack.pop())
	fmt.Println(intStack.isEmpty())
	fmt.Println(intStack)

	var sStack = stack[string]{}
	sStack.push("hey")
	sStack.push("hello")
	sStack.push("how are you")

}
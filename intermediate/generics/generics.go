package main
import "fmt"

type stack[T any] struct{
	elements []T
}

func (s *stack[t]) push(element t) {
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


func main(){
	var stack = stack[int]{}
	stack.push(20)
	stack.push(50)
	stack.push(100)

	fmt.Println(stack.pop())
	fmt.Println(stack.isEmpty())

}
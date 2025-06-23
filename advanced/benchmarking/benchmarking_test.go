// we do `go text "filename"`
// for testing files
package main

import (
	"fmt"
	"testing"
)

func add(a, b int) int {
	return  a+b
}


func TestAdd(t *testing.T){
	result := add(2,3)
	expected := 6

	if result != expected{
		t.Errorf("add(2, 3) = %d want %d", result, expected)
	}
}

func main(){
	fmt.Println("hello world")

}
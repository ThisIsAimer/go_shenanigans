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



//even in error it executes all arguments
func TestAddTableDriven(t *testing.T){
	tests := []struct{a,b,expected int}{
		{2,3,5},
		{2,5,8},
		{5,3,8},
		{2,8,9},
	}

	for _, test :=  range tests{
		result := add(test.a,test.b)
		if result != test.expected{
			t.Errorf("add(%d,%d) = %d but we want %d",test.a,test.b,result,test.expected)
		}
	}
}


func TestAddSubTests(t *testing.T){
	tests := []struct{a,b,expected int}{
		{2,3,5},
		{5,9,8},
		{5,3,8},
		{9,8,9},
	}

	for _, test :=  range tests{
		
		t.Run(fmt.Sprintf("add(%d,%d)",test.a,test.b),func(t *testing.T) {
			result := add(test.a,test.b)
			if result!= test.expected{
				t.Errorf("result is: %d, expected: %d",result,test.expected)
			}
		})
	}
}


func main(){
	fmt.Println("hello world")

}
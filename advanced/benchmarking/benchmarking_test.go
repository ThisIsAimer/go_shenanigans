package main

import (
	"fmt"
	"testing"
)

// for bechmarking
// go test -bench=. e:\coding\golang\basics\advanced\benchmarking\benchmarking_test.go|grep -v 'cpu:'

func add(a, b int) int {
	return  a+b
}

func BenchMarkAdd(b *testing.B){
 for range b.N{
	add(9,10)
 }
}

func main(){
	fmt.Println("hello world")

}
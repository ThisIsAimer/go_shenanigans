package main

import (
	"fmt"
	"reflect"
)
// manupulate at runtime
func main(){

	x := 77

	y := reflect.ValueOf(x)
	v := y.Type()

	fmt.Println("Value type:", v)

}
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

	fmt.Println("Value x type:", v)
	fmt.Println("Value x Kind:", v.Kind())
	fmt.Println("Value x is int?:", v.Kind() == reflect.Int)
	fmt.Println("Value x is string?:", v.Kind() == reflect.String)
	fmt.Println("is x zero?:", y.IsZero())

	fmt.Println("-----------------------------------------------")

	a := 70

	v1 := reflect.ValueOf(&a).Elem()
	v2 := reflect.ValueOf(&a)
	fmt.Println("v1 is type:", v1.Type())
	fmt.Println("v2 is type:", v2.Type())

	fmt.Println("v1 int value is:", v1.Int())
	v1.SetInt(20)
	fmt.Println("now, v1 int value is:", v1.Int())

	fmt.Println("-----------------------------------------------")

	var itf any = "Golang"

	myReflect := reflect.ValueOf(itf)

	fmt.Println("type of itf:", myReflect.Type())

	if( myReflect.Kind() == reflect.String){
		fmt.Println("itf is of string type")
	}
	

	


}
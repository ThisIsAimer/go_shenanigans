package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

type greeter struct{}

func (g greeter) Greet(name string) string {
	return "Hello " + name
}

func (g greeter) Intro(name string, name1 string) string {
	return "My name is " + name + "nice to meet you " + name1
}

// manupulate at runtime
func main() {

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

	if myReflect.Kind() == reflect.String {
		fmt.Println("itf is of string type")
	}

	fmt.Println("-----------------------------------------------")

	p := person{Name: "Gangotri", Age: 21}

	val := reflect.ValueOf(p)

	for i := range val.NumField() {
		fmt.Printf("at index %d, value %v\n", i, val.Field(i))
	}

	//manipulating the struct with reflect
	val1 := reflect.ValueOf(&p).Elem()

	nameField := val1.FieldByName("Name")

	if nameField.CanSet() {
		nameField.SetString("Godavari")
	} else {
		fmt.Println("cant change person")
	}

	fmt.Println("-----------------------------------------------")

	fmt.Println("person is:", p)

	g := greeter{}

	greet := reflect.TypeOf(g)

	greetValue := reflect.ValueOf(g)

	// var meth reflect.Method

	fmt.Println("type:", greet)

	for i := range greet.NumMethod() {
		method := greet.Method(i)

		//method at 0 index is the greet method
		fmt.Printf("method no. %d: %v\n", i, method.Name)
	}

	myMethod := greetValue.MethodByName("Intro")

	result := myMethod.Call([]reflect.Value{reflect.ValueOf("Shia"),reflect.ValueOf("Sunni")})
	// []string{"Friend"}
	// []type{type("sjfaopj"), type("jkasebfj")}
	fmt.Println("greetings results:", result[0].String())

}

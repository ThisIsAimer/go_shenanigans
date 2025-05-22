package main

import (
	"fmt"
	"math"
)


type Geomerty interface{
	area() float64
	perim() float64
}

type Rect struct{
	hight, width float64
}

type Circle struct{
	radius float64
}


func (r Rect) area() float64{
	return r.hight * r.width
}

func (c Circle) area() float64{
	return math.Pi * c.radius *c.radius
}

func (r Rect) perim() float64{
	return 2*( r.hight + r.width)
}

func (c Circle) perim() float64{
	return 2*math.Pi*c.radius
}

func (c Circle) diametre() float64{
	return 2*c.radius
}

func measure(g Geomerty) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){

	var rectangle = Rect{hight: 10,width: 15}
	var circle = Circle{radius: 10}
	measure(rectangle)
	measure(circle)

	myPrinter("hello",10,20.5,false)
	evaluate(true)
}


//interface type takes any value
func myPrinter ( a ...interface{}){
	for _,v := range a{
		fmt.Println(v)
	}
}
//or
// func myPrinter ( a ...any){
// 	for _,v := range a{
// 		fmt.Println(v)
// 	}
// }

func evaluate(a any) {
	switch a.(type) {
		case int: fmt.Println("its int")
		case string: fmt.Println("its string")

		default: fmt.Println("its not defined")
	}
}
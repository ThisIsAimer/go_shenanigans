package main
 
import (
    "fmt"
)
 
type Speaker interface {
    Speak() string
}
 
type Dog struct{ breed string}


//for god to be of speaker interface, it must have this func 
func (d Dog) Speak() string {
    return "Woof!"
}
 
func main() {
    // var s Speaker
    // d := Dog{ breed: "german shep"}
	// s=d
	// fmt.Println(s)
    // fmt.Println(s == d)

	var s Speaker = Dog{breed: "german shep"}
    d := Dog{ breed: "german shep"}
	fmt.Println(s)
    fmt.Println(s == d)
}
package main
import "fmt"

func main(){
	type person struct{
		firstName string
		lastName string
		age int

	}

	var person1 = person{
		firstName: "rudra",
		lastName: "shiva",
		age: 12,
	}

	person2 := person{
		firstName: "gothum",
		age: 19,
	}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person2)

}
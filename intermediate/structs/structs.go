package main

import (
	"fmt"
)

type person struct{
		firstName string;
		lastName string;
		age int;

	}

func main(){

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
	fmt.Println("----------------------------------------------------------------")

	//this is a method
	fmt.Println(person1.fullName())

	//ananomous structs
	user := struct{
		username string
		email string
	}{
		username: "ichigo",
		email: "whatever@whatever.com",
	}

	fmt.Println(user)

}

//methods

func (p person) fullName() string{
	return fmt.Sprintln(p.firstName,p.lastName)
}
package main

import (
	"fmt"
)

type Person struct{
		firstName string;
		lastName string;
		age int;
		address Address;
		Phone

	}

type Phone struct{
	number int
}

type Address struct{
	city string
	country string
}

func main(){

	var person1 = Person{
		firstName: "Rudra",
		lastName: "Shiva",
		age: 14,
		address: Address{
			city: "Uttarakhand",
			country: "Bharat",
		},
		Phone: Phone{
			number: 23464646,
		},
	}

	person2 := Person{
		firstName: "Gothum",
		age: 19,
	}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	fmt.Println(person1.address.city)
	person2.lastName = "Aanand"
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

	//mutable struct with pointer methods
	fmt.Println("person age before",person1.age)
	person1.birthday()
	fmt.Println("person age after",person1.age)

	fmt.Println(person1.number)//ananomous sturucts can be used like this


}

//methods

func (p Person) fullName() string{
	return fmt.Sprintln(p.firstName,p.lastName)
}
func (p *Person) birthday() {
	p.age++
}
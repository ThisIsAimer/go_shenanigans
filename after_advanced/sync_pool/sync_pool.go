package main

import (
	"fmt"
	"sync"
)

type person struct{
	name string
	age int
}

type levels struct{
 level int
}

func main(){

	var pool = sync.Pool{
		New: func()any{
			fmt.Println("creating a new person")
			return &person{}
		},
	}

	// the .(*person) is telling get that i expect the value returned by pool.Get() to be of type *person
	// type assertion only works when it returns an interface(any) type
	person1 := pool.Get().(*person)
	person1.name = "Robin"
	person1.age = 10

	fmt.Println("we got person 1:",person1)

	fmt.Println("---------------------------------------------------")
	// put puts back the instance in queue
	pool.Put(person1)
	fmt.Println("putting back person 1")

	person2 := pool.Get().(*person)
	fmt.Println("person 2:", person2)
	fmt.Println("---------------------------------------------------")

	person3 := pool.Get().(*person)
	fmt.Println("person 3:", person3)
	person3.name = "Pookie"
	person3.age = 18

	pool.Put(person2)
	pool.Put(person3)
	fmt.Println("putting both persons back")

	fmt.Println("---------------------------------------------------")
	person4 := pool.Get().(*person)
	fmt.Println("person 4:", person4)

	person5 := pool.Get().(*person)
	fmt.Println("person 5:", person5)
	// works like queue

	person6 := pool.Get().(*person)
	fmt.Println("person 6:", person6)

	//we can put anything in existing pool
	pool.Put(&levels{level: 99})

	fmt.Println("---------------------------------------------------")
	Mylevel := pool.Get().(*levels)
	fmt.Println("my level is:", Mylevel)

	//default value of pool will be created by the pool's new value
	person7 := pool.Get().(*person)
	fmt.Println("person 7:", person7)



}


// we create a new instance only when there is no data stored through put value
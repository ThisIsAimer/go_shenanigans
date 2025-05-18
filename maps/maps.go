package main

import (
	"fmt"
	"maps"
)

func main(){
	// maps in golang
	// var maps map[int]string
	// maps := make(map[int]string)
	var my_map = map[int]string{
		1 : "hey",
		2: "friend",
	}
	my_map[3] = "bro"

	fmt.Println(my_map)

	// delete(maps,2)

	fmt.Println(my_map[2], my_map[1])

	// clear(maps)
	// fmt.Println(maps)

	//present checks if the value exists
	var value, ok = my_map[1]
	fmt.Println(value)
	fmt.Println("Value exists with key[1]?:",ok)

	map1 := map[int]string{
		1 : "as",
		2 : "ok",
	}

	fmt.Println("maps equal?:", maps.Equal(map1,my_map))
	fmt.Println("maps equal?:", "ok"=="ok")

	//a map made with var map1 map[string]int is a nil map
	//and we cant put any value in it
	//to initialise a empty map in which we can add values,
	//we use the maka function
	// var foods map[int]int
	// foods[1]=1
	// fmt.Println(foods)
	//does work


	var marks = make(map[string]int)
	fmt.Println(marks)

	marks["akriti"] = 100

	fmt.Println(marks)

	map_in_map := make(map[int]map[int]string)
	map_in_map[1] = my_map
	map_in_map[2] = map1
	fmt.Println(map_in_map)

}
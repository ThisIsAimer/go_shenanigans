package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

type byAge []Person

type byName []Person

func (a byAge) Len() int {
	return len(a)
}

func (a byAge) Less(i, j int) bool {
	return a[i].age < a[j].age
}

func (a byAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a byName) Len() int {
	return len(a)
}

func (a byName) Less(i, j int) bool {
	return a[i].name < a[j].name
}

func (a byName) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {

	numbers := []int{56, 573, 24, 26, 8, 3, 1353, 36, 36, 91}
	sort.Ints(numbers)
	fmt.Println("sorted numbers", numbers)

	myStr := []string{"hello", "algorithm", "bye", "golang", "Friend", "Aimer", "Andro", "Sundrago", "string"}
	sort.Strings(myStr)
	fmt.Println("sorted strings", myStr)

	fmt.Println("--------------------------------------------------------")

	people := []Person{
		{"Bishakha", 20},
		{"Astra", 25},
		{"Supreet", 20},
		{"Andro", 19},
		{"Friend", 30},
		{"Rudra", 37},
		{"Aimer", 21},
	}

	fmt.Println("unsorted:\n", people)

	sort.Sort(byAge(people))
	fmt.Println("sorted people by age:\n", people)

	sort.Sort(byName(people))
	fmt.Println("sorted people by name:\n", people)

	fmt.Println("--------------------------------------------------------")
}

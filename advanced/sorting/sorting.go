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
//--------------------------------------------------------------------------

type by func (a,b *Person) bool

type PersonSorter struct{
	People []Person
	by func (a,b *Person) bool
}

func(a *PersonSorter) Len() int{
	return len(a.People)
}

func (a *PersonSorter) Swap(i, j int) {
	a.People[i], a.People[j] = a.People[j], a.People[i]
}

func (a *PersonSorter) Less(i,j int) bool{
	return a.by(&a.People[i],&a.People[j])
}

func (by by) sort(people []Person) {
	ps := &PersonSorter{
		People: people,
		by: by,
	}

	sort.Sort(ps)
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
	
	age := func(p1,p2 *Person) bool{
		return p1.age < p2.age
	}

	by(age).sort(people)
	fmt.Println("people again sorted by age", people)

	fmt.Println("--------------------------------------------------------")


	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guava"}
	
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted by last character:", stringSlice)

}

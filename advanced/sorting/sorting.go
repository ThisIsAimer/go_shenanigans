package main

import (
	"fmt"
	"sort"
)

func main(){

	numbers := [] int {56,573,24,26,8,3,1353,36,36,91}
	sort.Ints(numbers)
	fmt.Println("sorted numbers", numbers)

	myStr := [] string {"hello","algorithm","bye","golang","Friend","Aimer","Andro","Sundrago","string"}
	sort.Strings(myStr)
	fmt.Println("sorted strings", myStr)

}
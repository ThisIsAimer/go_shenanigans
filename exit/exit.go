package main

import (
	"fmt"
	"os"
)



func main(){
	defer fmt.Println("defer hey")
	//defer doesnt execute on os.Exit like panic

	fmt.Println("before exit")

	os.Exit(1)
	// non zero exit code means error

	fmt.Println("hey there")

}
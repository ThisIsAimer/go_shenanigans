package main

import (
	"fmt"
	"os"
)

func main(){

	os.Setenv("USER", "Bro")
	os.Setenv("HOME", "Multiverse")

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("user:", user)
	fmt.Println("home:", home)

}
package main
import "fmt"

//the init function is used to initialise packages
//and are executed once before executing main

func init(){
	fmt.Println("initialising package 1")
}

func init(){
	fmt.Println("initialising package 2")
}

func init(){
	fmt.Println("initialising package 3")
}



func main(){
	fmt.Println("in the main function now")

}
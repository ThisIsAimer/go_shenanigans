package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){

	fmt.Println("command:", os.Args[0])

	//go run "e:\coding\golang\basics\intermediate\cmd_line_args\cmd_line_args.go" hello bro are are you
	for index, args := range os.Args{
		fmt.Printf("argument %d: %s\n", index, args)
	}

	//defining flags
	//go run "e:\coding\golang\basics\intermediate\cmd_line_args\cmd_line_args.go" -name "androEz bro" -age 19 -male 1
	////go run "e:\coding\golang\basics\intermediate\cmd_line_args\cmd_line_args.go" -name androEz -age 19
	var name string
	var age int
	//if we pass male, it will be set to true!
	var male bool

	flag.StringVar(&name,"name","bro","your name")
	flag.IntVar(&age,"age", 18, "your age")
	flag.BoolVar(&male,"male",false,"are you male?")

	flag.Parse()

	fmt.Println("name:",name)
	fmt.Println("age:", age)
	fmt.Println("male?:", male)

}
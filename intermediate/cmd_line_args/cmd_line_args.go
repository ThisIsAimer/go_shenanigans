package main

import (
	"fmt"
	"os"
)

func main(){

	fmt.Println("command:", os.Args[0])

	//go run "e:\coding\golang\basics\intermediate\cmd_line_args\cmd_line_args.go" hello bro are are you
	for index, args := range os.Args{
		fmt.Printf("argument %d: %s\n", index, args)
	}

}
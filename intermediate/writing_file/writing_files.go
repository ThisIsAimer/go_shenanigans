package main

import (
	"fmt"
	"os"
)

func main(){
	newFile, err := os.Create("Output.txt")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("hello world")
	defer newFile.Close()

	byteData := []byte("I love golang its soo much fun!\nI think its an amazing language!\n")


	// file will be written in the directory we are in
	num , err := newFile.Write(byteData)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	println("bytes written:", num)

}
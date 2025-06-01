package main

import (
	"fmt"
	"os"
)

func main(){
	newFile, err := os.Create("output.txt")
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

	strFile, err := os.Create("strOutput.txt")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	defer strFile.Close()

	var strData = "One piece is real!!!!!!\nA proud member of the strawhat crew.\n"

	num, err = strFile.WriteString(strData)

	println("strings written:", num)

}
package main

import (
	"fmt"
	"os"
)

func main(){
	//it doesnt create directory
	newFile, err := os.Create("basics/intermediate/writing_file/output.txt")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	// we use
	//err = os.MkdirAll("intermediate/writing_file", os.ModePerm)
	// we can use 0755 instead of os.ModePerm
	// to make directory

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

	strFile, err := os.Create("basics/intermediate/writing_file/strOutput.txt")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	defer strFile.Close()

	var strData = "One piece is real!!!!!!\nA proud member of the strawhat crew.\n"

	num, err = strFile.WriteString(strData)

	println("strings written:", num)

}
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main(){
	//buffering is reading or writing huge amounts of data in chunks
	//its used in video processing, youtube and communication

	//strings.new reader is a reader but not buffered
	var strReader = strings.NewReader("hello go bufferio package\n")
	var reader = bufio.NewReader(strReader)

	//it will contain the read string
	data := make([]byte,20)
	fmt.Println()
	//returns number of bytes it reads
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("the data read is %d and the data is: %s\n", n,data[:n])

	line, err := reader.ReadString('\n')//must use single quotes

	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	//reader give out the remaining string that was left after transfering
	//string into data.
	println("string read is:", line)

}
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
	var reader = strings.NewReader("hello go bufferio package\n")
	var bufReader = bufio.NewReader(reader)

	//it will contain the read string
	data := make([]byte,20)
	fmt.Println()
	//returns number of bytes it reads
	n, err := bufReader.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("the data read is %d and the data is: %s\n", n,data[:n])

}
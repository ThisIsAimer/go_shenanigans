package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Println("-------------------------------------------------------")

	var write = bufio.NewWriter(os.Stdout)

	wData := []byte("hello how are you doing!")
	aData := []byte("\nI love golang")

	//written in an internal buffer
	n, err = write.Write(wData)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	n, err = write.Write(aData)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	
	var myString = "\nhey whats up"

	n, err = write.WriteString(myString)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("bytes written:", n)

	fmt.Println("\nwritten string is:-")
	fmt.Println("-------------------------------------------------------")
	err = write.Flush()
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	

}
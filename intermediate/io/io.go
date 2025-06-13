package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func readFromReader (reader io.Reader){
	buffer := make([]byte,1024)
	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Println("error is:", err)
		log.Fatal(err)
	}

	fmt.Println(string(buffer[:n]))
}

func writeToWriter (writer io.Writer, data string){
	_, err := writer.Write([]byte(data))
	if err != nil {
		fmt.Println("error is:", err)
		log.Fatal(err)
	}

}


func closeResource (resource io.Closer){
	err := resource.Close()
	if err != nil {
		fmt.Println("error is:", err)
		log.Fatal(err)
	}
}

func bufferExample(c string){
	var buffer bytes.Buffer //memory on stack
	_, err := buffer.WriteString(c)

	if err != nil {
	   log.Fatal(err)
	}
}


func multibufferExample (c ...string){
	var readers []io.Reader
	for _,element := range c{
		readers = append(readers, strings.NewReader(element))
	}

	multireader := io.MultiReader(readers...)

	buffer := new(bytes.Buffer) //allocated memory on heap

	_, err := buffer.ReadFrom(multireader)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println(buffer.String())
	
}

func main(){
	bufferExample("Hey friends!")
	multibufferExample("I am Great!\n","how are you?\n","I love golang\n")

}
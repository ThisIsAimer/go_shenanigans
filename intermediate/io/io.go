package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
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

func pipeExample(word string){
	reader, writer := io.Pipe()

	go func ()  {
		writer.Write([]byte(word))
		writer.Close()
	}()// goroutine will be transferred to the next thread and following lines of code
		// will be executed normally

	buffer := new(bytes.Buffer)

	buffer.ReadFrom(reader)

	fmt.Println(buffer.String())
}

func main(){
	bufferExample("Hey friends!")
	multibufferExample("I am Great!\n","how are you?\n","I love golang\n")


	fmt.Println("Reading from reader!")
	//reading from reader
	readFromReader(strings.NewReader("this is reading the string"))
	//---------------------------------------------------------

	file, err := os.ReadFile(`intermediate/io/golang.txt`)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	readFromReader(strings.NewReader(string(file)))

	//or

	myfile, err := os.OpenFile(`intermediate/io/golang.txt`, os.O_APPEND, 0755) //with flags
	//myfile statisfies the io.reader interface wth its read function
	
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	defer closeResource(myfile)
	fmt.Println("------------------------------------------------------------------")
	readFromReader(myfile)

	var writer bytes.Buffer
	writeToWriter(myfile,"writing to writer\n")

	myfile, err = os.OpenFile(`intermediate/io/golang.txt`, os.O_APPEND, 0755)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	byteSL := make([]byte,1204)
	n, err := myfile.Read(byteSL)

	writeToWriter(&writer, string(byteSL[:n]))

	fmt.Println("writer contains:",writer.String())

	// piping------------------------------------------

	pipeExample("this is the pute thing")

}
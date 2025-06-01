package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("intermediate/reading_file/backend.txt")
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("file closed")
	}()

	var data = make([]byte, 895)

	_, err = file.Read(data)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("file content is:", string(data))

	fmt.Println("-----------------------------------------------")

	file1, err := os.Open("intermediate/reading_file/backend.txt")
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	defer func() {
		file1.Close()
		fmt.Println("file1 closed")
	}()

	myBuffer := bufio.NewScanner(file1)

	//scans each individual line of the file and puts it in my buffer
	for myBuffer.Scan() {
		line := myBuffer.Text()
		//assigns the string read from file to line

		fmt.Println(line)

	}

	err = myBuffer.Err()
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

}

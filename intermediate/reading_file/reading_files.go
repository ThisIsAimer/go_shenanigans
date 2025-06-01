package main

import (
	"fmt"
	"os"
)

func main(){

	file, err :=os.Open("intermediate/reading_file/backend.txt")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	defer func(){
		file.Close()
		fmt.Println("file closed")
	}()


	var data = make([]byte,890)

	_, err = file.Read(data)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}


	fmt.Println("file content is:", string(data))

}
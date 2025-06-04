package main

import (
	"fmt"
	"os"
)

func main(){
	var tempFile, err = os.CreateTemp("intermediate/temp_files","temp_file")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	
	defer func(){
		tempFile.Close()
		err = os.Remove(tempFile.Name())
		if err != nil {
			fmt.Println("error is:",err)
			return
		}
	}()

	fmt.Println("temp file created:",tempFile.Name())

}
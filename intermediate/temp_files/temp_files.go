package main

import (
	"fmt"
	"os"
)

func main(){
	//var tempFile, err = os.CreateTemp("","temp_file")
	//empty string in first argument will store the temp file in operating systems temp file folder


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

	fmt.Println("temporary file:",tempFile.Name())

	tempDir, err := os.MkdirTemp("intermediate/temp_files","temp_dir")
	fmt.Println("temporary directory:",tempDir) 

	defer func(){
		err = os.RemoveAll(tempDir)
		if err != nil {
			fmt.Println("error is:",err)
			return
		}
	}()
}
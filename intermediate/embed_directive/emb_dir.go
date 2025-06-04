package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//must be in the same folder

//go:embed example.txt
var content string

//go:embed golang
var embDir embed.FS

func main() {

	fmt.Println("embeded content:", content)
	
	fileContent, err := embDir.ReadFile("golang/love.txt")
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	fmt.Println("love.txt:", string(fileContent))

	fileContent, err = embDir.ReadFile("golang/golang.txt")
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	fmt.Println("-------------------------------------------------")
	fmt.Println("golang.txt:\n", string(fileContent))
	fmt.Println("-------------------------------------------------")

	//this takes fs.FS unline filepath.walkdir that takes string
	err = fs.WalkDir(embDir,"golang",func(path string, d fs.DirEntry, err error) error{
		if err != nil {
			fmt.Println("error is:",err)
			return err
		}

		fmt.Println("file is:",path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}


}
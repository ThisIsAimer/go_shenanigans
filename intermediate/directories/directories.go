package main

import (
	"fmt"
	"os"
)

func main(){
	err := os.MkdirAll("./intermediate/directories/subdir",0755)
	checkError(err)

	//defer os.RemoveAll("./intermediate/directories/subdir")

	err = os.WriteFile("./intermediate/directories/subdir/golang.txt",[]byte("golang is sooo good!"),0755)
	checkError(err)

	err = os.MkdirAll("./intermediate/directories/c/a/b",0755)
	checkError(err)

	dir, err := os.ReadDir("./intermediate/directories")
	checkError(err)

	for _, element := range dir{
		fmt.Println(element)
		fmt.Println(element.IsDir())
		fmt.Println("----------------------------------------------------------")
	}

	fmt.Println("----------------------------------------------------------")

	//creates directory the base
	checkError(os.Chdir("intermediate/directories"))
	dir, err = os.ReadDir(".")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("reading intermediate/directories: ")

	fmt.Println(dir)

}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}
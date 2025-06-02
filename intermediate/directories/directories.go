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


	fmt.Println("hello world")

}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}
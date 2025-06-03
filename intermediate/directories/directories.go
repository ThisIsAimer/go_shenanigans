package main

import (
	"fmt"
	"os"
	"path/filepath"
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
	checkError(os.Chdir("./intermediate/directories"))
	//checkError(os.Chdir("intermediate/directories"))
	curDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	
	fmt.Println("current directory:",curDir)

	dir, err = os.ReadDir(".")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("reading intermediate/directories: ")

	fmt.Println("read directory: ",dir)

	checkError(os.Chdir("./../.."))
	curDir, err = os.Getwd()
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	
	fmt.Println("current directory:",curDir)

	fmt.Println("----------------------------------------------------------")

	//filepath.Walk and filepath.Walkdir 
	//walkdir is better

	fmt.Println("walking directory: ")
	var pathFile = "./intermediate/directories"
	err = filepath.WalkDir(pathFile,func(path string, d os.DirEntry, err error) error{
		if err != nil {
			fmt.Println("error is:",err)
			return err
		}
		fmt.Println("current path is: ",path)
		return nil
	})

	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	/*
	func myWalker(path string, d os.DirEntry, err error) error {
    if err != nil {
        fmt.Println("error is:", err)
        return err
    }
    fmt.Println(path)
    return nil
	}

	filepath.WalkDir(pathFile, myWalker)
	*/

	defer func(){
		err = os.RemoveAll("./intermediate/directories/subdir")
		if err != nil {
			fmt.Println("error is:",err)
			return
		}

		err = os.RemoveAll("./intermediate/directories/c")
		if err != nil {
			fmt.Println("error is:",err)
			return
		}
	}()

	// os.Remove to remove just one entity

}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}
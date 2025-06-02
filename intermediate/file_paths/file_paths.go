package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {

	reletivePath := "./docs/go.txt"
	//as i am in windows
	var absolutePath = "C:/home/users/docs/file.go"

	var joinedPath = filepath.Join("home", "downloads", "golang.txt")
	fmt.Println("file path:", joinedPath)

	normalisePath := filepath.Clean("./data/../data/golang.txt")

	fmt.Println("cleaned parth:", normalisePath)

	var dir, file = filepath.Split("home/docs/golang.go")

	fmt.Printf("directorypath is: %s, and filename is: %s\n",dir,file)
	fmt.Println("base:",filepath.Base("home/docs/golang.txt"))

	fmt.Println("absolute path?:",filepath.IsAbs(absolutePath))
	fmt.Println("absolute path?:",filepath.IsAbs(reletivePath))


	fmt.Println("--------------------------------------------------")
	fmt.Println("extention of absolute path:",filepath.Ext(absolutePath))
	fmt.Println("fine name without suff:",strings.TrimSuffix(file,filepath.Ext(file)))
	//fmt.Println("fine name without suff:",strings.TrimSuffix(file,".go"))

	file, err := filepath.Rel("a/b","a/b/c/d/golang.go")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("reletive path:", file)


	//goes back 2 folders
	file, err = filepath.Rel("a/t/y","a/b/c/d/golang.go")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("reletive path:", file)

	//takes cd in terminal as reference
	absPath,err := filepath.Abs(reletivePath)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	fmt.Println("absolute path:",absPath)

}

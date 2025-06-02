package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var joinedPath = filepath.Join("home", "downloads", "golang.txt")
	fmt.Println("file path:", joinedPath)

	normalisePath := filepath.Clean("./data/../data/golang.txt")

	fmt.Println("cleaned parth:", normalisePath)

}

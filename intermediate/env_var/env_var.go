package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	err := os.Setenv("USER", "Bro")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	err = os.Setenv("HOME", "Multiverse")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	user := os.Getenv("USER")
	home := os.Getenv("HOME")

	fmt.Println("user:", user)
	fmt.Println("home:", home)

	//
	for _, pair := range os.Environ() {
		kvPair := strings.SplitN(pair,"=",2)
		fmt.Println(kvPair[0]) //1 for value
	}

	err = os.Unsetenv("HOME")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	err = os.Unsetenv("USER")
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Println("user:", os.Getenv("USER"))
	fmt.Println("home:", os.Getenv("HOME"))

	fmt.Println("unset env var")

	fmt.Println("-----------------------------------------------------")

	str := "a=b=c=d"

	fmt.Println("-1:", strings.SplitN(str,"=",-1))
	fmt.Println(" 0:", strings.SplitN(str,"=",0))
	fmt.Println(" 1:", strings.SplitN(str,"=",1))
	fmt.Println(" 2:", strings.SplitN(str,"=",2))
	fmt.Println(" 3:", strings.SplitN(str,"=",3))
	fmt.Println(" 4:", strings.SplitN(str,"=",4))
	fmt.Println(" split:", strings.Split(str,"="))

}

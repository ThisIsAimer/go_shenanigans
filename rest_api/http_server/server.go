package main

import (
	"fmt"
	"net/http"
)

func main(){

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// writes to resp
		fmt.Fprintln(resp, "hello server!")
	})

	// 127.0.0.1 is for local host and 3000 post is our gate number
	const serverAdd string = "127.0.0.1:3000"


	fmt.Println("listening to port 3000")

	err := http.ListenAndServe(serverAdd,nil)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("hello world")

}

// we go to:
// http://localhost:3000/
// to see the response
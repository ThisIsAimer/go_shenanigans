package main

import (
	"fmt"
	"net/http"
)

func main(){

	// "/" is base root for http://localhost:3000/
	// w is resp and r is req as golang conventions
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// writes to resp
		fmt.Fprintln(w, "Golang is soo much fun!")
	})

	// 127.0.0.1 is for local host and 3000 post is our gate number
	// golang assumes that its local host if we just add the port with :
	const port string = ":3000"


	fmt.Println("listening to port,", port)

	err := http.ListenAndServe(port,nil)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("hello world")

}

// we go to:
// http://localhost:3000/
// to see the response
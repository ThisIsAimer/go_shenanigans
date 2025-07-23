package main

import (
	"fmt"
	"net/http"
	"strings"
)

type user struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Place string `json:"place"`
}



func homeRoute(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "string")
	fmt.Println("someone accessed: home")
	fmt.Println("method:", r.Method)

	switch r.Method{
	case http.MethodGet:
		fmt.Fprintln(w, "accessed : Home. with: Get")
		// for a specific ID, it will be routed to /teachers/90 or smth
		urlPath := strings.TrimPrefix(r.URL.Path, "/teachers/")
		userId := strings.TrimSuffix(urlPath,"/")

		if userId != ""{
			fmt.Println("id is:",userId)
		}

		//------------handling quary-----------------------------------
		if len(r.URL.Query()) > 0{

			quaryParams := r.URL.Query()

			for key := range quaryParams{
				fmt.Printf("key %v : %v \n", key, quaryParams.Get(key))
			}
		}

		
	case http.MethodPost:
		fmt.Fprintln(w, "accessed : Home. with: Post")
	case http.MethodPut:
		fmt.Fprintln(w, "accessed : Home. with: Put")
	case http.MethodPatch:
		fmt.Fprintln(w, "accessed : Home. with: Patch")
	case http.MethodDelete:
		fmt.Fprintln(w, "accessed : Home. with: Delete")
	default:
		fmt.Fprintln(w, "accessed : Home")

	}

}

func main(){
	http.HandleFunc("/", homeRoute)

	port := 3000

	server := &http.Server{
		Addr: fmt.Sprintf(":%d",port),
		Handler: nil,
	}

	fmt.Println("server is running on port:", port)


	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

}
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main(){
	client := &http.Client{}

	// mock apis
	
	// resp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")
	// resp, err := client.Get("https://catfact.ninja/fact")
	// resp, err := client.Get("https://dog.ceo/api/breeds/image/random")
	// resp, err := client.Get("https://reqres.in/api/users/2")
	// resp, err := client.Get("https://api.open-meteo.com/v1/forecast?latitude=35&longitude=139&current_weather=true")

	resp, err := client.Get("https://api.genderize.io?name=sunpreet")


	if err != nil {
		fmt.Println("error making get request:", err)
		return
	}
	defer resp.Body.Close()

	// reading and printing response

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body:", err)
		return
	}

	fmt.Println("body is:", string(body))

}
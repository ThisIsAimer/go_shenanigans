package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	myUser := User{Name: "Rudra", Email: "rudra@example.com"}

	jsonData, err := json.Marshal(myUser)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("created json data:")
	fmt.Println(string(jsonData))

	fmt.Println("----------------------------------------------------------------")


	customJsonData :=  `{"name":"Draupadi","email":"draupadi@example.com"}`

	var customUser User

	err = json.Unmarshal([]byte(customJsonData), &customUser)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	println("our custon struct:", customUser.Name ,customUser.Email)

}

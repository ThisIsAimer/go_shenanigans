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
	user := User{Name: "Rudra", Email: "rudra@example.com"}

	jsonData, err := json.Marshal(user)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("created json data:\n", string(jsonData))

}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
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

	fmt.Println("----------------------------------------------------------------")

	decodableJsonData := `{"name":"Song Jin Woo","email":"soloLeveling@example.com"}`

	reader := strings.NewReader(decodableJsonData)
	decoder := json.NewDecoder(reader)

	var decodedUser User

	decoder.Decode(&decodedUser)

	fmt.Println("decoded User:")
	fmt.Println(decodedUser)


	fmt.Println("----------------------------------------------------------------")

	encodableUser := User{Name: "Wukong", Email: "backMyth@china.com"}

	var buf bytes.Buffer

	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(encodableUser)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}


	fmt.Println("encoded json data:")
	fmt.Print(buf.String())

}

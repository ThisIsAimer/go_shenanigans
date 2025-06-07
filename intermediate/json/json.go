package main

import (
	"encoding/json"
	"fmt"
)

// we are exporting these! so these will be caps
type Person struct {
	// json keeps keys as lowe case
	//we use the json tag for the json package

	//json package will look for the json tag
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	friend := Person{Name: "Friend", Age: 34}

	//marshaling
	//In Go, "marshal" means to convert data (like a struct or map) into a specific format
	jsonData, err := json.Marshal(friend)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	fmt.Println("json data:", string(jsonData))
}

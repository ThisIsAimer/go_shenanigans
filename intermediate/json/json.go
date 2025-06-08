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
	Name    string  `json:"name"`
	Age     int     `json:"age,omitempty"` //will omit if not provided
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	friend := Person{Age: 34, Email: "friend@gmail.com", Address: Address{City: "Patna", State: "Bihar"}}
	rudra := Person{Name: "Rudra", Email: "rudra@gmail.com"}

	//marshaling
	//In Go, "marshal" means to convert data (like a struct or map) into a specific format
	jsonData, err := json.Marshal(friend)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	fmt.Println("json data:", string(jsonData))

	jsonData1, err := json.Marshal(rudra)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}
	fmt.Println("json data:", string(jsonData1))

	var myEmployee Employee

	myData := `{"full_name":"Akrit", "employee_id":"060204","age": 21,"address":{"city":"Rishikesh","state":"Uttarakhand"}}`

	err = json.Unmarshal([]byte(myData), &myEmployee)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("employee is:", myEmployee)
	fmt.Println("age by 2032:", myEmployee.Age+(2032-2025))

	addresses := []Address{
		{City: "Mumbai", State: "Maharashtra"},
		{City: "Bengaluru", State: "Karnataka"},
		{City: "Chennai", State: "Tamil Nadu"},
		{City: "Kolkata", State: "West Bengal"},
		{City: "Hyderabad", State: "Telangana"},
	}

	addressList, err := json.Marshal(addresses)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("addresses are", string(addressList))

	// handling unknown json structures

	jsonPerson := `{ "name": "Priya Singh","age": 25,"email": "priya.singh@example.com","address": {"city": "Lucknow","state": "Uttar Pradesh" }}`

	var data map[string]any

	err = json.Unmarshal([]byte(jsonPerson), &data)
	if err != nil {
		fmt.Println("mapping error is:", err)
		return
	}

	fmt.Println("------------------------------------------------------------------------")


	fmt.Println("mapped data:", data)
	fmt.Println("mapped data:", data["address"])
	fmt.Println("mapped data:", data["name"])
	fmt.Println("mapped data:", data["age"])
	

}

type Employee struct {
	FullName   string  `json:"full_name"`
	EmployeeID string  `json:"employee_id"`
	Age        int     `json:"age"`
	Address    Address `json:"address"`
}

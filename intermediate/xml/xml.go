package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	//XMLName is case sensetive
	XMLName xml.Name `xml:"person"` //root email
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

func main() {
	//Extensible Markup Language
	//used for encoding encoding docs which is both human and machine readable

	fmt.Println("hello world")

	myFriend := Person{Name: "Sunpreet Singh", Age: 30, City: "random", Email: "email@example.com"}

	xmlData, err := xml.Marshal(myFriend)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("xml data is:", string(xmlData))

	xmlData, err = xml.MarshalIndent(myFriend, "", " ")

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("indented xml data is:", string(xmlData))

	fmt.Println("----------------------------------------------------")

	rawXmlData := `<person><name>Akriti</name><age>21</age><email>akriti@example.com</email></person>`
	var AkData Person

	err = xml.Unmarshal([]byte(rawXmlData), &AkData)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("unmarshaled datais:", AkData)
}

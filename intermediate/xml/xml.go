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
	Address Address  `xml:"address,omitempty"` //structs cant be omitemptied
	Email   string   `xml:"email"`
}

type Address struct {
	City  string `xml:"city,omitempty"`
	State string `xml:"state,omitempty"`
}

func main() {
	//Extensible Markup Language
	//used for encoding encoding docs which is both human and machine readable

	fmt.Println("hello world")

	myFriend := Person{Name: "Sunpreet Singh", Age: 30, Email: "email@example.com"}
	myFriend1 := Person{Name: "Vikas", Age: 30, Email: "email@example.com", Address: Address{City: "New Delhi", State: "Union Terratory"}}

	xmlData, err := xml.Marshal(myFriend)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("sunpreet xml data is:", string(xmlData))

	xmlData, err = xml.MarshalIndent(myFriend1, "", " ")

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("vikas indented xml data is:", string(xmlData))

	xmlData, err = xml.MarshalIndent(myFriend, "", " ")

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("sunpreet indented xml data is:", string(xmlData))

	fmt.Println("----------------------------------------------------")

	rawXmlData := `<person><name>Akriti</name><age>21</age><email>akriti@example.com</email></person>`
	var AkData Person

	err = xml.Unmarshal([]byte(rawXmlData), &AkData)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("unmarshaled datais:", AkData)

	updateXml := `<person><name>Srekanth</name><age>21</age><email>srekanth@example.com</email><address><city>Raipur</city><state>chattisgarh</state></address></person>`

	var randomPerson Person

	err = xml.Unmarshal([]byte(updateXml), &randomPerson)

	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("Srekanth's xml is:", randomPerson)

	//----------------------------------------
	fmt.Println("Local:", randomPerson.XMLName.Local)
	fmt.Println("Space:", randomPerson.XMLName.Space)
	//attributes

	book := Book{ISBN: "545-68574-4575457",Title: "Love", Auther: "Me", Psudo: "psudo", PsudoAttr:"attr" }

	xmlBook, err:= xml.MarshalIndent(book,""," ")
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("attributes of book:", string(xmlBook))

}

type Book struct {
	XMLName xml.Name `xml:"book"`
	//attr means attribute
	ISBN    string   `xml:"isbn,attr,omitempty"`
	Title   string   `xml:"title,attr"`
	Auther  string   `xml:"auther,attr"`
	Psudo string  `xml:"psudo"`
	PsudoAttr string  `xml:"psudo,attr"`
}

package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// base 64 is binary to text encoding
	// it converts binary data into textual representation
	// it uses a set of 64 ascii characters
	// its used to transmit binary data over text based protocals like email
	// and also databases and urls
	// uses uppercase, lowercase, 0-9, +, /,
	// = for padding at the end of encodded data

	// in go we use encoding/base64 package

	var data = []byte("hey i love golang!")

	var encoded = base64.StdEncoding.EncodeToString(data)

	fmt.Println("encoded data:", encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error is:", err)
		return
	}

	fmt.Println("------------------------------------------------------")

	fmt.Println("decoded byte:", decoded)
	//bytes can be directly converted to strings
	fmt.Println("text form:", string(decoded))
	//fmt.Printf("Text format: %s\n",decoded)

}

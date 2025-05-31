package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	// hashing is a process in computing that transforms data
	// into a fixed size string of characters
	// it typically appears random
	// implimented by a hash function
	// if we use SHA-256 hash it makes a 256 bit or 32 byte long hash
	// SHA-512 creates a 512 bit or 64 byte long hash
	// hashing generates an ir-reversable output
	// its hard to easily reverse an hash
	// sha-256 can create 2^256 possible hash values

	var password = "password1234"
	var hashed = sha256.Sum256([]byte(password))
	var hashed1 = sha512.Sum512([]byte(password))
	fmt.Println("hashed value is:", hashed)

	//must be printed with %x
	fmt.Printf("hashed sha-256 value is: %x\n", string(hashed[:]))
	fmt.Printf("hashed sha-512 value is: %x\n", string(hashed1[:]))

}

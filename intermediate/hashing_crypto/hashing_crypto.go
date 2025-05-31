package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
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

	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")

	salt,err := generateSalt()
	if err != nil {
		fmt.Println("error is:",err)
		return
	}
	fmt.Printf("salt is: %x\n",salt)

	var signupPass = hashPassward(password,salt)

	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("base 64 encoded salt:",encodedSalt) // stored to database
	fmt.Println("Hash:",signupPass) // stored to database

	fmt.Println("------------------------------------------------------")
	fmt.Println("------------------------------------------------------")

	//encoded salt is retrieved from database
	decodedSalt,err := base64.StdEncoding.DecodeString(encodedSalt)
	if err != nil {
		fmt.Println("error is:",err)
		return
	}

	fmt.Printf("base 64 decoded salt is:%x\n",decodedSalt)

	loginPass := hashPassward(password,decodedSalt)

	if loginPass == signupPass{
		fmt.Println("logged in successfully!")
	} else{
		fmt.Println("login failed, please check username and password")
	}


	//what we are doing here is generating a salt and storing it to database after
	//base64 encoding it.

	//we make a new hash by passing the passward and the original salt value
	//hashing it and then again base64 encoding it. we store it in database

	//for login
	//we get the base64 encoded salt then decoding it to get the original salt
	
	//after that we hash the passward with the original salt to get a hash that
	//will be exactly same as the hash stord in the database if passward is same.

	//then we compare both hashed values for login

}



func generateSalt() ([]byte,error){
	salt := make([]byte,16)
	_, err := io.ReadFull(rand.Reader,salt)
	if err != nil {
		return nil , err
	} else{
		return salt , nil
	}	
}


func hashPassward(pass string, salt []byte) string {
	saltedPassward := append(salt,[]byte(pass)...)

	hash := sha256.Sum256(saltedPassward)
	return base64.RawStdEncoding.EncodeToString(hash[:])
}

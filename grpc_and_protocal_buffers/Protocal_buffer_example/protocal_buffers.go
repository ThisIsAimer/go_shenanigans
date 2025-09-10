package main

import "fmt"

// generate proto files for golang
// protoc  --go_out=. ./proto/main.proto

// after importing
// protoc -I=proto  --go_out=. proto/main.proto proto/user/user.proto proto/order.proto


//packages compiling into same folder must have same package name 
// /proto/gen;mainpb
// /proto/gen/user;userpb

func main() {
	fmt.Println("hello world")

}

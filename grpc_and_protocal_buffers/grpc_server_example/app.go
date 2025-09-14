package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

func main() {

	port := ":50051"

	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer()

	fmt.Println("server running at port " + port)
	err = grpcServer.Serve(lis)

	if err != nil {
		fmt.Println("couldnt serve:", err)
		return
	}

}

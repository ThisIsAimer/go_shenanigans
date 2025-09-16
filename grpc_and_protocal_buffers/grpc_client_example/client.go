package main

import (
	"context"
	"fmt"
	pb "grpcclient/proto/gen"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	cert := `certificate\certificate.pem`

	cred, err := credentials.NewClientTLSFromFile(cert, "")
	if err != nil {
		fmt.Println("error setting client:", err)
		return
	}

	port := ":50051"

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(cred))
	if err != nil {
		fmt.Println("error setting client:", err)
		return
	}

	defer conn.Close()

	client := pb.NewCalculateClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.AddRequest{
		A: 30,
		B: 40,
	}

	res, err := client.Add(ctx, &req)
	if err != nil {
		fmt.Println("error getting result:", err)
		return
	}

	fmt.Println("sum result:", res.Sum)

}

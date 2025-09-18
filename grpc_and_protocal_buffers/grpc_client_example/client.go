package main

import (
	"context"
	"fmt"
	pb "grpcclient/proto/gen"
	hugpb "grpcclient/proto/gen/hugs"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	// for add service

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

	state := conn.GetState().String()

	fmt.Println("state is:", state)

	// for greet service

	client2 := pb.NewGreeterClient(conn)

	reqGreet := pb.HelloRequest{Message: "hello how are you!"}

	res2, err := client2.Greet(ctx, &reqGreet)

	if err != nil {
		fmt.Println("error getting response:", err)
	}
	fmt.Println("greet response:", res2.Message)

	//get hugging service
	client3 := pb.NewHuggingClient(conn)

	hugReq := hugpb.HugRequest{Name: "Andro"}

	res3, err := client3.GiveHug(ctx, &hugReq)

	if err != nil {
		fmt.Println("error getting hug:", err)
	}

	fmt.Println("got hug:", res3.Message)

}

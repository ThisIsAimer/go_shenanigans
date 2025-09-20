package main

import (
	"context"
	"fmt"
	pb "grpcstreamclient/proto/gen"
	"io"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	port := ":50051"

	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error making client:", err)
		return
	}

	client := pb.NewCalculateClient(conn)

	ctx := context.Background()

	
	// server side streaming
	febReq := &pb.FibonacchiRequest{N: 20}

	stream, err := client.GenerateFibonacchi(ctx, febReq)
	if err != nil {
		fmt.Println("error calling GenerateFibonacchi:", err)
		return
	}

	for {
		resp, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("end of stream!")
			break
		} else if err != nil {
			fmt.Println("error recieving data from GenerateFibonacchi:", err)
			return
		}

		fmt.Println("got number:", resp.GetNumber())
	}

	// client side streaming

	stream1, err := client.SendNumbers(ctx)
	if err != nil {
		fmt.Println("error creating stream:", err)
		return
	}

	for i := range 10 {
		err := stream1.Send(&pb.SendRequest{Number: int32(i)})
		if err != nil {
			fmt.Println("error is:", err)
			return
		}
		time.Sleep(time.Second)
	}

	resp1, err := stream1.CloseAndRecv()
	if err != nil {
		fmt.Println("error closing and recieving stream:", err)
		return
	}

	fmt.Println("sum is:", resp1.Sum)

}

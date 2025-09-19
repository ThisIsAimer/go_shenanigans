package main

import (
	"context"
	"fmt"
	pb "grpcstreamclient/proto/gen"
	"io"

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

	febReq := &pb.FibonacchiRequest{N: 20}

	strem, err := client.GenerateFibonacchi(ctx, febReq)
	if err != nil {
		fmt.Println("error calling GenerateFibonacchi:", err)
		return
	}

	for {
		resp, err := strem.Recv()

		if err == io.EOF {
			fmt.Println("end of stream!")
			break
		} else if err != nil {
			fmt.Println("error recieving data from GenerateFibonacchi:", err)
			return
		}
		
		fmt.Println("got number:", resp.GetNumber())
	}

}

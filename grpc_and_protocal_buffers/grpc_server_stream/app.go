package main

import (
	"context"
	"fmt"
	pb "grpcstream/proto/gen"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	sum := req.A + req.B

	return &pb.AddResponse{
		Result: sum,
	}, nil

}


// new function style for streaming
func (s *server) GenerateFibonacchi(req *pb.FibonacchiRequest, stream pb.Calculate_GenerateFibonacchiServer) error {
	n := req.N

	a, b := 0, 1

	for range int(n) {
		err := stream.Send(&pb.FibonacchiResponse{Number: int32(a)})
		if err != nil {
			return err
		}
		a, b = b, a+b
		time.Sleep(time.Second)
	}

	return nil
}

func main() {
	port := `:50051`

	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("error listening to server:", err)
		return
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculateServer(grpcServer, &server{})

	fmt.Println("listening to port:", port)

	err = grpcServer.Serve(lis)

	if err != nil {
		fmt.Println("error serving:", err)
		return
	}

}

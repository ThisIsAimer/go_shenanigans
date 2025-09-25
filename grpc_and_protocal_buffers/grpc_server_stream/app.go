package main

import (
	"bufio"
	"context"
	"fmt"
	pb "grpcstream/proto/gen"
	"io"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("no metadata")
	} else {
		fmt.Println("metadata:", md)
	}

	va, ok := md["authorization"]

	if !ok {
		fmt.Println("no authorization header")
	} else {
		fmt.Println("authorization:", va)
	}

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
	}

	return nil
}

// for client side streaming
func (s *server) SendNumbers(stream pb.Calculate_SendNumbersServer) error {

	var sum int32

	for {

		req, err := stream.Recv()

		if err == io.EOF {
			fmt.Println("end of stream")
			return stream.SendAndClose(&pb.SendResponse{Sum: sum})
		} else if err != nil {
			fmt.Println("error getting request:", err)
			return err
		}

		println(req.GetNumber())

		sum += req.GetNumber()
	}

}

//omni directional streaming

func (s *server) Chat(stream pb.Calculate_ChatServer) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		//reciving value
		req, err := stream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("error text:", err)
			return err
		}

		fmt.Println("------------------------------------------------------")
		fmt.Println("recieved Text:", req.GetText())
		//read input from terminal

		fmt.Print("please end next text with a \".\": ")
		input, err := reader.ReadString('.')
		if err != nil {
			return err
		}

		input = strings.TrimSpace(input)

		// sending data
		err = stream.Send(&pb.ChatMessage{Text: input})
		if err != nil {
			return err
		}

	}

	fmt.Println("returning control")
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

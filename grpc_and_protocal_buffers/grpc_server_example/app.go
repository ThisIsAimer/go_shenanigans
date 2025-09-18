package main

import (
	"context"
	"fmt"
	"net"

	pb "grpcapp/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct {
	pb.UnimplementedCalculateServer
	pb.UnimplementedGreeterServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	sum := req.A + req.B

	return &pb.AddResponse{
		Sum: sum,
	}, nil

}

func (s *server) Greet(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	request := req.Message

	return &pb.HelloResponse{
		Message: fmt.Sprintf("Got: %s\ngreetings!", request),
	}, nil

}

func main() {

	cert := `certificate\certificate.pem`
	key := `certificate\key.pem`

	port := ":50051"

	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	creds, err := credentials.NewServerTLSFromFile(cert, key)
	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterCalculateServer(grpcServer, &server{})
	pb.RegisterGreeterServer(grpcServer, &server{})

	fmt.Println("server running at port " + port)
	err = grpcServer.Serve(lis)

	if err != nil {
		fmt.Println("couldnt serve:", err)
		return
	}

}

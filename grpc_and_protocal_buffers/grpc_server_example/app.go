package main

import (
	"context"
	"fmt"
	"net"

	pb "grpcapp/proto/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {

	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil

}

func main() {

	port := ":50051"

	lis, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println("failed to listen:", err)
		return
	}

	grpcServer := grpc.NewServer()

	pb.RegisterCalculateServer(grpcServer, &server{})

	fmt.Println("server running at port " + port)
	err = grpcServer.Serve(lis)

	if err != nil {
		fmt.Println("couldnt serve:", err)
		return
	}

}

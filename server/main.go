package main

import (
	"context"
	"fmt"
	"grpc-practice/pb"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesResponse, error) {
	fmt.Println("ListFiles was invoked with %v", req)

	// read storage directory
	paths, err := ioutil.ReadDir("./storage")
	if err != nil {
		return nil, err
	}
	var filenames []string
	for _, path := range paths {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}
	res := &pb.ListFilesResponse{
		Filenames: filenames,
	}
	return res, nil

}

func main() {
	lis, err := net.Listen("tcp", "localhost:50052")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})

	fmt.Println("Server is running on port 50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

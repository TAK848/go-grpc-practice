package main

import (
	"context"
	"fmt"
	"grpc-practice/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)
	callListFiles(client)

}

func callListFiles(client pb.FileServiceClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalf("error while calling ListFiles RPC: %v", err)
	}
	fmt.Println(res.GetFilenames())
}

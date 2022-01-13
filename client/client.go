package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc-sample/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 建立gRPC連線
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewHelloServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// 向gRPRC server發送請求
	resp, err := grpcClient.Hello(ctx, &pb.HelloReq{Name: name})
	if err != nil {
		log.Fatalf("could not get response: %v", err)
	}
	log.Printf("Message: %s", resp.GetMessage())
}

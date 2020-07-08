package main

import (
	"context"
	"fmt"
	pb "github.com/adriananeci/grpc-sample-app/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	args := os.Args

	server:=getEnv("GRPC_SERVER", "127.0.0.1")
	serverPort :=getEnv("GRPC_PORT", "5300")

	conn, err := grpc.Dial(server + ":" +serverPort, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewReverseClient(conn)
	request := &pb.Request{
		Message: args[1],
	}
	response, err := client.Do(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	fmt.Println(response.Message)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
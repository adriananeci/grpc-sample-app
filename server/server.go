package main

import (
	pb "github.com/adriananeci/grpc-sample-app/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5300")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterReverseServer(grpcServer, &server{})
	grpcServer.Serve(listener)
}

type server struct{}

func (s *server) Do(c context.Context, request *pb.Request) (response *pb.Response, err error) {
	n := 0
	rune := make([]rune, len(request.Message))
	// create an array of runes to safely reverse a string
	for _, r := range request.Message {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse using runes
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	response = &pb.Response{
		Message: output,
	}
	return response, nil
}
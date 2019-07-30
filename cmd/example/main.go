package main

import (
	"net"

	"google.golang.org/grpc"

	pb "github.com/dio/grpctranscoding/api/example/v1"
	"github.com/dio/grpctranscoding/pkg/example"
)

func main() {
	listener, _ := net.Listen("tcp", ":8000")

	var serverOptions = []grpc.ServerOption{}
	grpcServer := grpc.NewServer(serverOptions...)

	pb.RegisterExampleServer(grpcServer, example.NewService())
	grpcServer.Serve(listener)
}

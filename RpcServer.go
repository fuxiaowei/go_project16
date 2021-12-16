package main

import (
	"go_project16/proto"
	"go_project16/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	rpcServer := grpc.NewServer()
	proto.RegisterPersonServiceServer(rpcServer, &service.Server{})
	// Register reflection service on gRPC server.
	reflection.Register(rpcServer)
	if err := rpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

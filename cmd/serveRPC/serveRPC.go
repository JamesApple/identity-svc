package main

import (
	"fmt"
	"net"
	"root"

	"root/rpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	server := &rpc.Server{}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8081))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	root.RegisterIdentityServer(grpcServer, server)
	reflection.Register(grpcServer)

	grpcServer.Serve(listener)
}

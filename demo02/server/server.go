package main

import (
	"google.golang.org/grpc"
	"grpcstudy/demo02/server/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}

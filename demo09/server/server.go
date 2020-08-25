package main

import (
	"google.golang.org/grpc"
	"grpcstudy/demo09/server/helper"
	"grpcstudy/demo09/server/services"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(lis)
}

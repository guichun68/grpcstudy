package main

import (
	"google.golang.org/grpc"
	"grpcstudy/demo07/server/services"
	"log"
	"net"
)

func main() {

	//rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, err := net.Listen("tcp",":8099")
	if err != nil{
		log.Fatal(err)
	}
	rpcServer.Serve(lis)
}

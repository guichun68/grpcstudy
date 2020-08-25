package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcstudy/demo4/server/services"
	"log"
	"net"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("keys/server.crt","keys/server_no_passwd.key")
	if err != nil{
		log.Fatal(err)
	}
	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}

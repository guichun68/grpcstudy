package main

import (
	"google.golang.org/grpc"
	"grpcstudy/demo19/server/helper"
	"grpcstudy/demo19/server/services"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCreds()))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))//商品服务
	services.RegisterOrderServiceServer(rpcServer, new(services.OrdersService))//订单服务
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))//用户服务
	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(lis)
}

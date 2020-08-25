package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcstudy/demo8/client/helper"
	services "grpcstudy/demo8/client/service"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient := services.NewProdServiceClient(conn)

	prodRes2 , err := prodClient.GetProdStocks(context.Background(),&services.QuerySize{Size: 3})
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("库存列表获取结果:",prodRes2)
	fmt.Println(prodRes2.Prodres[2])
	fmt.Println(prodRes2.Prodres[2].ProdStock)
}

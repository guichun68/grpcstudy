package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	services "grpcstudy/demo03/client/service"
	"log"
)

func main() {
	conn, err:=grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient:=services.NewProdServiceClient(conn)
	prodRes, err := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId:20})
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
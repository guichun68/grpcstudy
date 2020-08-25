package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcstudy/demo10/client/helper"
	. "grpcstudy/demo10/client/service"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient := NewProdServiceClient(conn)
	//获取ProductId=12的商品在区域C的库存数
	prodRes, err := prodClient.GetProdInfo(context.Background(), &ProdRequest{ProdId: 12,})
	fmt.Println(prodRes)
	fmt.Println("prod_name:",prodRes.ProdName)
}

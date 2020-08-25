package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"
	"grpcstudy/demo11/client/helper"
	. "grpcstudy/demo11/client/service"
	"log"
	"time"
)

func main() {

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	orderClient := NewOrderServiceClient(conn)
	timeGo := time.Now()
	timeProto, err := ptypes.TimestampProto(timeGo)
	if err != nil {
		fmt.Println(err)
	}
	orderClient.NewOrder(context.Background(), &OrderMain{OrderId: 31456, OrderNo: "bj009", UserId: 2048, OrderMoney: 349.90, OrderTime: timeProto})
}

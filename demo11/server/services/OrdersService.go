package services

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"log"
)

type OrdersService struct {
}

func (this *OrdersService) NewOrder(ctx context.Context, orderMain *OrderMain) (*OrderResponse, error) {
	fmt.Println(orderMain)
	orderTime, err := ptypes.Timestamp(orderMain.OrderTime)
	if err != nil{
		log.Fatal(err)
	}
	timeGo:=orderTime.Format("2006-01-02 15:04:05")
	fmt.Println("orderTime:",orderTime)
	fmt.Println("goTime:",timeGo)
	fmt.Printf("server收到订单信息:订单号:%v, 订单金额:%v,订单时间:%v",orderMain.OrderNo,orderMain.OrderMoney, timeGo)

	return &OrderResponse{Status: "OK", Message: "success",}, nil
}

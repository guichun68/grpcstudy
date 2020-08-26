package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcstudy/demo17/client/helper"
	. "grpcstudy/demo17/client/service"
	"io"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	userClient := NewUserServiceClient(conn)
	ctx := context.Background()

	var i int32
	req := UserScoreRequest{}
	req.Users = make([]*UserInfo, 0)
	for i = 1; i < 6; i++ {
		req.Users = append(req.Users, &UserInfo{UserId: i})
	}
	stream, err := userClient.GetUserScoreByServerStream(ctx, &req)
	if err!=nil{
		log.Fatal(err)
	}
	for{
		res, err := stream.Recv()
		if err == io.EOF{
			break
		}
		if err != nil{
			log.Fatal(err)
		}
		fmt.Println(res.Users)
	}
}

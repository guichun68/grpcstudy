package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcstudy/demo19/client/helper"
	. "grpcstudy/demo19/client/service"
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

	stream, err := userClient.GetUserScoreByTWS(ctx)
	if err != nil{
		log.Fatal(err)
	}

	var uid int32 = 1
	for j := 1; j <= 3; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i = 1; i <= 5; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: uid})
			uid ++
		}
		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}
		res, err := stream.Recv()
		if err == io.EOF{
			log.Println(err)
		}
		fmt.Println(res.Users)
	}
/*	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)*/
}

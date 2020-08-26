package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcstudy/demo18/client/helper"
	. "grpcstudy/demo18/client/service"
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

	stream, err := userClient.GetUserScoreByClientStream(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for j := 1; j <= 3; j++ {
		req := UserScoreRequest{}
		req.Users = make([]*UserInfo, 0)
		for i = 1; i <= 5; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: i})
		}
		err := stream.Send(&req)
		if err != nil {
			log.Println(err)
		}
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

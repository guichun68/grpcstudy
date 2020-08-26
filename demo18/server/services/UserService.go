package services

import (
	"context"
	"io"
	"time"
)

type UserService struct {
}

//普通方法
func (*UserService) GetUserScore(ctx context.Context, in *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{Users: users}, nil
}

//服务端流式响应
func (*UserService) GetUserScoreByServerStream(in *UserScoreRequest, stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			users = (users)[0:0]
		}
		time.Sleep(1 * time.Second)
	}
	if len(users) > 0 {
		return stream.Send(&UserScoreResponse{Users: users})
	}
	return nil
}

func (*UserService) GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err := stream.Recv()
		if err == io.EOF { //接收完毕
			return stream.SendAndClose(&UserScoreResponse{Users: users})
		}
		if err != nil {
			return err
		}
		for _, user := range req.Users {
			user.UserScore = score //这里好比是服务端做的业务处理
			score++
			users = append(users, user)
		}
	}
}

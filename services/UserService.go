package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

type UserService struct {

}

// 普通方法
func (this *UserService) GetUserScore(ctx context.Context,request  *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for _, user := range request.Users {
		user.UserScore = score
		score++
		users = append(users,user)
	}
	return &UserScoreResponse{Users:users}, nil
}

// 服务端流
func (this *UserService) GetUserScoreByServerStream(request *UserScoreRequest,stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo,0)
	for index, user := range request.Users {
		user.UserScore = score
		score++
		users = append(users,user)
		// 每隔两条发送
		if (index + 1) % 2 == 0 && index > 0 {
			//log.Printf("send data %v", users)
			err := stream.Send(&UserScoreResponse{Users:users})
			if err != nil {
				fmt.Println(err)
			}
			//发送完数据后，将users切片强制清空
			users = (users)[0:0]
		}
		time.Sleep(time.Second * 1)

	}
	// 剩余users数据
	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users:users})
		if err != nil {
			fmt.Println(err)
		}
	}
	return  nil
}

//客户端流
func (this *UserService) GetUserScoreByClientStream(stream UserService_GetUserScoreByClientStreamServer) error  {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err :=  stream.Recv()
		if err == io.EOF { //接收客户端信息完成
			return stream.SendAndClose(&UserScoreResponse{Users:users})
		}
		if err != nil {
			return err
		}
		for _, user := range req.Users  { //服务端业务处理
			user.UserScore = score
			score++
			users = append(users, user)
		}
	}
}

//双向流

func (this *UserService) GetUserScoreByTWS(stream UserService_GetUserScoreByTWSServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for {
		req, err :=  stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		for _, user := range req.Users  { //服务端业务处理
			user.UserScore = score
			score++
			users = append(users, user)
		}
		err = stream.Send(&UserScoreResponse{Users:users})
		if err != nil {
			log.Println(err)
		}
		//发送完数据后，将users切片强制清空,实际开发中不这么使用
		users = (users)[0:0]
	}
}
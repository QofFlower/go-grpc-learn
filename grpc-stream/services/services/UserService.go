package services

import (
	"context"
	"io"
	"log"
	"math/rand"
)

type UserService struct {
}

func (s *UserService) GetUserScore(cxt context.Context, userRequest *UserScoreRequest) (*UserScoreResponse, error) {
	res := make([]*UserInfo, 0)
	for _, user := range userRequest.Users {
		user.UserScore = rand.Int31()
		res = append(res, user)
	}
	return &UserScoreResponse{Users: res}, nil
}

// 服务端流模式返回数据
func (s *UserService) GetUserScoreByServerStream(userRequest *UserScoreRequest, serverStream UserService_GetUserScoreByServerStreamServer) error {
	res := make([]*UserInfo, 0)
	for index, user := range userRequest.Users {
		user.UserScore = rand.Int31n(300)
		res = append(res, user)
		if (index+1)%2 == 0 { // 没隔两条发送一次
			err := serverStream.Send(&UserScoreResponse{Users: res})
			if err != nil {
				return err
			}
			res = (res)[0:0]
		}
	}
	if len(res) > 0 {
		err := serverStream.Send(&UserScoreResponse{Users: res})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *UserService) GetUserScoreByClientStream(clientStream UserService_GetUserScoreByClientStreamServer) error {
	res := make([]*UserInfo, 0)
	for {
		usr, err := clientStream.Recv()
		if err == io.EOF {
			return clientStream.SendAndClose(&UserScoreResponse{Users: res})
		}
		if err != nil {
			return err
		}
		for _, user := range usr.Users {
			user.UserScore = rand.Int31n(300)
			res = append(res, user)
		}
	}
}

func (s *UserService) GetUserScoreByDoubleEndStream(stream UserService_GetUserScoreByDoubleEndStreamServer) error {
	res := make([]*UserInfo, 0)
	for {
		usr, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
			return nil
		}
		for _, user := range usr.Users {
			user.UserScore = rand.Int31n(100)
			res = append(res, user)
		}
		err2 := stream.Send(&UserScoreResponse{Users: res})
		if err2 != nil {
			log.Fatal(err2)
			return err2
		}
	}
}

func (s *UserService) mustEmbedUnimplementedUserServiceServer() {
}

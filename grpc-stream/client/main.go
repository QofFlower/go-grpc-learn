package main

import (
	"context"
	"fmt"
	"go-rpc/grpc-first/client/helper"
	. "go-rpc/grpc-first/client/services"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(helper.GetClientCredentials()))

	if err != nil {
		log.Fatal("Failed to connect the server", err)
	}
	defer func(cc *grpc.ClientConn) {
		err := cc.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(cc)

	cxt := context.Background()
	usc := NewUserServiceClient(cc)

	stream, err3 := usc.GetUserScoreByDoubleEndStream(cxt)
	if err3 != nil {
		log.Fatal(err3)
		return
	}

	for j := 0; j < 3; j++ {
		req := &UserScoreRequest{}
		for i := 0; i < 5; i++ {
			req.Users = append(req.Users, &UserInfo{UserId: int32(i + 1)})
		}
		err2 := stream.Send(req)
		if err2 != nil {
			log.Fatal(err2)
			break
		}
		res, err4 := stream.Recv()
		if err4 == io.EOF {
			break
		}
		if err4 != nil {
			log.Fatal(err4)
		}
		fmt.Println(res.Users)
	}
}

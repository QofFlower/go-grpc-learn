package main

import (
	"context"
	"fmt"
	"go-rpc/grpc-first/client/helper"
	"go-rpc/grpc-first/client/services"
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

	prodClient := services.NewProdServiceClient(cc)
	cxt := context.Background()

	res, err := prodClient.GetProdStocks(cxt, &services.QuerySize{Size: 10})
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(res.Productes)
}

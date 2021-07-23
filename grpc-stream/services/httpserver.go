package main

import (
	"context"
	"go-rpc/helper"
	"go-rpc/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	endPoint := "localhost:10086"
	cxt := context.Background()

	err := services.RegisterProdServiceHandlerFromEndpoint(cxt, gwmux, endPoint, opt)
	if err != nil {
		log.Fatal(err)
	}

	err1 := services.RegisterOrderServiceHandlerFromEndpoint(cxt, gwmux, endPoint, opt)
	if err1 != nil {
		log.Fatal(err1)
	}

	httpServer := &http.Server{
		Addr:    ":9468",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}

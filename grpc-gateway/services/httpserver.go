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

	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:10086", opt)

	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":9468",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}

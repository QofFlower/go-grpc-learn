package main

import (
	"go-rpc/helper"
	"go-rpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredential()))
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})

	l, err := net.Listen("tcp", ":10086")
	if err != nil {
		log.Fatal(err)
		return
	}
	rpcServer.Serve(l)
}

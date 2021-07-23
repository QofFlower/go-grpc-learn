package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"go-rpc/grpc-first/client/services"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	cert, err := tls.LoadX509KeyPair("cert/client.pem", "cert/client.key")
	if err != nil {
		log.Fatal("Failed to load the credentials key of client", err)
	}
	certPool := x509.NewCertPool()
	ca, err1 := ioutil.ReadFile("cert/ca.pem")
	if err1 != nil {
		log.Fatal("Failed to load the public key", err1)
	}
	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	cc, err := grpc.Dial(":10086", grpc.WithTransportCredentials(creds))
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
	prodRes, err1 := prodClient.GetProdStock(context.Background(), &services.ProdRequest{ProdId: 10086})
	if err1 != nil {
		log.Fatal("Encountering error during requestint to the server", err1)
	}
	fmt.Println(prodRes.ProdStock)
}

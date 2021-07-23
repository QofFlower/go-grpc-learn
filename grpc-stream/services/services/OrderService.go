package services

import (
	"context"
	"fmt"
)

type OrderService struct {
}

func (s *OrderService) NewOrder(cxt context.Context, orderRequest *OrderRequest) (*OrderResponse, error) {
	fmt.Println(orderRequest)
	return &OrderResponse{Status: "200", Message: "OK"}, nil
}

func (s *OrderService) mustEmbedUnimplementedOrderServiceServer() {
}

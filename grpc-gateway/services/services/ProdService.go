package services

import (
	context "context"
)

type ProdService struct {
}

func (t *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 114514}, nil
}

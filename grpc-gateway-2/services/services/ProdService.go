package services

import (
	context "context"
	"fmt"
)

type ProdService struct {
}

func (t *ProdService) GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 114514}, nil
}

func (t *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	return &ProdResponseList{Productes: []*ProdResponse{
		{ProdStock: 114514},
		{ProdStock: 10086},
		{ProdStock: 721},
		{ProdStock: 666},
	}}, nil
}

func (t *ProdService) mustEmbedUnimplementedProdServiceServer() {
	fmt.Println("Fuck you")
}

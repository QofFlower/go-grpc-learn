package services

import (
	context "context"
	"fmt"
)

type ProdService struct {
}

func (t *ProdService) GetProdStock(cxt context.Context, request *ProdRequest) (*ProdResponse, error) {
	var stock int32
	if request.ProdArea == ProdAreas_A {
		stock = 114514
	} else if request.ProdArea == ProdAreas_B {
		stock = 721
	} else if request.ProdArea == ProdAreas_C {
		stock = 666
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (t *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	return &ProdResponseList{Productes: []*ProdResponse{
		{ProdStock: 114514},
		{ProdStock: 10086},
		{ProdStock: 721},
		{ProdStock: 666},
	}}, nil
}

func (t *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	fmt.Println(in)
	return &ProdModel{ProdId: 22, ProdName: "GTX 3070"}, nil
}

func (t *ProdService) mustEmbedUnimplementedProdServiceServer() {
	fmt.Println("Fuck you")
}

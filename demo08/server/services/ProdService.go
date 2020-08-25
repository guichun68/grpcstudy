package services

import (
	"context"
	//"google.golang.org/grpc"
)

type ProdService struct {
}

//服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	Prodres := []*ProdResponse{
		&ProdResponse{ProdStock: 12},
		&ProdResponse{ProdStock: 13},
		&ProdResponse{ProdStock: 14},
	}
	return &ProdResponseList{Prodres: Prodres,}, nil
}

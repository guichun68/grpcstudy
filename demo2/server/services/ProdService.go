package services

import (
	"context"
	//"google.golang.org/grpc"
)
type ProdService struct {

}

func(this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest)(*ProdResponse, error){
	return &ProdResponse{ProdStock:20},nil
}
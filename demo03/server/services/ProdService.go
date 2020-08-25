package services

import (
	"context"
	//"google.golang.org/grpc"
)
type ProdService struct {

}
//服务具体实现
func(this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest)(*ProdResponse, error){
	return &ProdResponse{ProdStock:20},nil
}
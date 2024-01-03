package service

import (
	"context"
	"fmt"

	pb "github.com/kvii/issue3141/api/helloworld/v1"
	"github.com/kvii/issue3141/internal/biz"
)

type MyServiceService struct {
	pb.UnimplementedMyServiceServer
}

func NewMyServiceService(*biz.GreeterUsecase) *MyServiceService {
	return &MyServiceService{}
}

func (s *MyServiceService) Call(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println(req)
	return &pb.Response{}, nil
}

package service

import (
	"context"
	pb "github.com/go-zepto/templates/default/proto/app"
	"github.com/go-zepto/zepto"
)

type Service interface {
	Zepto() *zepto.Zepto
	Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error)
}

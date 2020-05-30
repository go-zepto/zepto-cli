package handlers

import (
	"context"
	pb "github.com/go-zepto/templates/default/proto/app"
	"github.com/go-zepto/templates/default/service"
)

type GRPCHandler struct {
	svc service.Service
}

func NewGRPCHandler(s service.Service) *GRPCHandler {
	return &GRPCHandler{
		svc: s,
	}
}

func (h GRPCHandler) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return h.svc.Hello(ctx, req)
}

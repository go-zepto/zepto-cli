package service

import (
	"context"
	pb "github.com/go-zepto/templates/default/proto/app"
	"github.com/go-zepto/zepto"
)


type svc struct {
	z *zepto.Zepto
}

func NewTodosService(z * zepto.Zepto) Service {
	return &svc{
		z: z,
	}
}

// Zepto - For convenience, to have access to the zepto instance
func (s *svc) Zepto() *zepto.Zepto {
	return s.z
}

func (s *svc) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Result: "Hello, " + req.Text,
	}, nil
}

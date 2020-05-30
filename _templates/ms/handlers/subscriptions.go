package handlers

import (
	"context"
	"github.com/go-zepto/zepto"
	pb "github.com/go-zepto/zepto-examples/todosvc/proto/todos"
	"github.com/go-zepto/zepto-examples/todosvc/service"
)

type Subs struct {
	svc service.Service
}

func NewSubsHandler(z * zepto.Zepto, s service.Service) *Subs {
	return &Subs{
		svc: s,
	}
}

func (s * Subs) TodoCreatedEvent(ctx context.Context, todo *pb.Todo) {
	s.svc.Zepto().Logger().Debug("[New event] todo.created: " + todo.Description)
}

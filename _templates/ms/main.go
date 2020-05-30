package main

import (
	"github.com/go-zepto/templates/default/handlers"
	pb "github.com/go-zepto/templates/default/proto/app"
	"github.com/go-zepto/templates/default/service"
	"github.com/go-zepto/zepto"
	"google.golang.org/grpc"
)

func main() {
	// Create Zepto
	z := zepto.NewZepto(
		zepto.Name("ms-hello"),
		zepto.Version("latest"),
	)

	svc := service.NewTodosService(z)

	// Setup gRPC Server
	z.SetupGRPC("0.0.0.0:9000", func(s *grpc.Server) {
		pb.RegisterHelloAppServer(s, handlers.NewGRPCHandler(svc))
	})

	// Setup HTTP Server
	z.SetupHTTP("0.0.0.0:8000", handlers.NewHTTPHandler(svc))

	// Setup Broker Provider (Google Pub/Sub)
	/*
		z.SetupBroker(gcp.NewBroker(
			gcp.ProjectID("my-project-id"),
			gcp.TopicPrefix("dev.hello."),
		))
	 */

	z.Start()
}

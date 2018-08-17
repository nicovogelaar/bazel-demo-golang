package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/nicovogelaar/bazel-demo-golang/pkg/logs"
	"google.golang.org/grpc"
	helloworldpb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	port = ":8080"
)

type greeterServer struct{}

func (s *greeterServer) SayHello(ctx context.Context, in *helloworldpb.HelloRequest) (*helloworldpb.HelloReply, error) {
	return &helloworldpb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	logs.Debug("starting greeter service")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	helloworldpb.RegisterGreeterServer(grpcServer, &greeterServer{})

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	logs.Debug("greeter service started")

	wg.Wait()
}

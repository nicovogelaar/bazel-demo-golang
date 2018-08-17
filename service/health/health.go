package main

import (
	"log"
	"net"
	"sync"

	"github.com/nicovogelaar/bazel-demo-golang/pkg/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

const (
	addr    = ":8081"
	service = "health"
)

func main() {
	logs.Debug("starting health service")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error setting up tcp listener: %v", err)
	}

	healthServer := health.NewServer()
	healthServer.SetServingStatus(service, healthpb.HealthCheckResponse_SERVING)

	server := grpc.NewServer()
	healthpb.RegisterHealthServer(server, healthServer)

	defer server.Stop()

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	logs.Debug("health service started")

	wg.Wait()
}

package handler

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func RegisterHealthCheck(server *grpc.Server) {
	grpc_health_v1.RegisterHealthServer(server, &healthCheck{})
}

type healthCheck struct {
}

func (h *healthCheck) Check(ctx context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (h *healthCheck) Watch(req *grpc_health_v1.HealthCheckRequest, stream grpc_health_v1.Health_WatchServer) error {
	return nil
}

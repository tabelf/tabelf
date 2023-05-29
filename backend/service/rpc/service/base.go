package service

import (
	"context"

	"google.golang.org/grpc/health/grpc_health_v1"
)

type BaseService struct {
	grpc_health_v1.UnimplementedHealthServer
}

// Check health check 接口.
func (s *BaseService) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

package service_test

import (
	"context"
	"tabelf/backend/service/rpc"
	"testing"

	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/health/grpc_health_v1"
	"tabelf/backend/service/rpc"
)

type BaseTestSuite struct {
	suite.Suite
	ctx context.Context
	rpc *rpc.GodServer
}

func (s *BaseTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.rpc = new(rpc.GodServer)
}

func (s *BaseTestSuite) Context() context.Context {
	return s.ctx
}

func (s *BaseTestSuite) TestCheck() {
	s.Run("test check", func() {
		resp, err := s.rpc.Check(s.ctx, &grpc_health_v1.HealthCheckRequest{})
		s.NoError(err)
		s.True(resp.Status == grpc_health_v1.HealthCheckResponse_SERVING)
	})
}

func (s *BaseTestSuite) TestGetTime() {
	s.Run("test get time", func() {})
}

func TestBaseTestSuite(t *testing.T) {
	suite.Run(t, new(BaseTestSuite))
}

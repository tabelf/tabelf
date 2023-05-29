package global_test

import (
	"context"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/testutils"
	"testing"

	service "tabelf/backend/service/api"
)

type SayHelloTestSuite struct {
	suite.Suite
	ctx    context.Context
	svc    *svc.ServiceContext
	router httpx.Router
	hp     testutils.Helper
}

func (s *SayHelloTestSuite) Context() context.Context {
	return s.ctx
}

func (s *SayHelloTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.svc, s.router = service.CreateTestApp()
	s.hp = testutils.Helper{Suite: s.Suite, Route: s.router, Config: s.svc.Config}
}

func TestSayHelloTestSuite(t *testing.T) {
	suite.Run(t, new(SayHelloTestSuite))
}

func (s *SayHelloTestSuite) TestSayHelloHandler() {
	s.Run("test say hello", func() {
		var result types.HelloResponse
		s.hp.Get("/api/hello/zhangsan").OK(&result)
		s.Equal("Hello zhangsan", result.Message)
	})
}

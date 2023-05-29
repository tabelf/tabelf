package global_test

import (
	"context"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/testutils"
	"testing"

	service "tabelf/backend/service/api"
)

type DatetimeTestSuite struct {
	suite.Suite
	ctx    context.Context
	svc    *svc.ServiceContext
	router httpx.Router
	hp     testutils.Helper
}

func (s *DatetimeTestSuite) Context() context.Context {
	return s.ctx
}

func (s *DatetimeTestSuite) SetupTest() {
	s.ctx = context.Background()
	s.svc, s.router = service.CreateTestApp()
	s.hp = testutils.Helper{Suite: s.Suite, Route: s.router, Config: s.svc.Config}
}

func TestDatetimeTestSuite(t *testing.T) {
	suite.Run(t, new(DatetimeTestSuite))
}

func (s *DatetimeTestSuite) TestGetDatetimeHandler() {
	s.Run("test get datetime", func() {
		s.hp.Get("/api/global/datetime").OK()
	})
}

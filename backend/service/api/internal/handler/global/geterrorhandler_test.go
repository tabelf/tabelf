package global_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/testutils"
	"testing"

	service "tabelf/backend/service/api"
)

type ErrorTestSuite struct {
	suite.Suite
	ctx    context.Context
	svc    *svc.ServiceContext
	router httpx.Router
	hp     testutils.Helper
}

func (s *ErrorTestSuite) Context() context.Context {
	return s.ctx
}

func (s *ErrorTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.svc, s.router = service.CreateTestApp()
	s.hp = testutils.Helper{Suite: s.Suite, Route: s.router, Config: s.svc.Config}
}

func TestErrorTestSuite(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}

func (s *ErrorTestSuite) TestGetErrorHandler() {
	s.Run("test 500", func() {
		r, err := http.NewRequestWithContext(s.ctx, http.MethodGet, "/api/global/error?code=500", nil)
		s.Nil(err)
		rr := httptest.NewRecorder()
		s.Panics(func() { s.router.ServeHTTP(rr, r) })
	})
	s.Run("test 400", func() {
		s.hp.Get("/api/global/error?code=400").Bad()
	})
}

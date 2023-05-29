package utils_test

import (
	"context"
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type NonceTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *NonceTestSuite) Context() context.Context {
	return s.ctx
}

func (s *NonceTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestNonceTestSuite(t *testing.T) {
	suite.Run(t, new(NonceTestSuite))
}

func (s *PemTestSuite) TestGenerateNonce() {
	_, err := utils.GenerateNonce()
	s.Nil(err)
}

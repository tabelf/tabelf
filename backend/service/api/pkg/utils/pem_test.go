package utils_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type PemTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *PemTestSuite) Context() context.Context {
	return s.ctx
}

func (s *PemTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestPemTestSuite(t *testing.T) {
	suite.Run(t, new(PemTestSuite))
}

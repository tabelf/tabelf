package utils_test

import (
	"context"
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type CompareTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *CompareTestSuite) Context() context.Context {
	return s.ctx
}

func (s *CompareTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestCompareTestSuite(t *testing.T) {
	suite.Run(t, new(CompareTestSuite))
}

func (s *CompareTestSuite) TestMaxInt() {
	s.Equal(3, utils.MaxInt([]int{1, 2, 3}...))
}

package utils_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type SliceTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *SliceTestSuite) Context() context.Context {
	return s.ctx
}

func (s *SliceTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestSliceTestSuite(t *testing.T) {
	suite.Run(t, new(SliceTestSuite))
}

func (s *SliceTestSuite) TestSliceContain() {
	s.Equal(true, utils.SliceContain([]string{"a", "b", "c"}, "a"))
	s.Equal(false, utils.SliceContain([]string{"a", "b", "c"}, "d"))
	s.Equal(true, utils.SliceContain([]string{""}, ""))
	s.Equal(false, utils.SliceContain([]string{}, ""))

	m := map[string]int{"a": 1}
	s.Equal(true, utils.SliceContain(m, "a"))
	s.Equal(false, utils.SliceContain(m, "b"))

	s.Equal(true, utils.SliceContain("abc", "a"))
	s.Equal(false, utils.SliceContain("abc", "d"))
}

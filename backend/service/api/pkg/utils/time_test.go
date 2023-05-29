package utils_test

import (
	"context"
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type TimeTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *TimeTestSuite) Context() context.Context {
	return s.ctx
}

func (s *TimeTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestTimeTestSuite(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}

func (s *TimeTestSuite) TestParseStartAndEndTime() {
	start := "2021-11-16"
	end := "2021-11-16"
	startTime, endTime, err := utils.ParseStartAndEndTime(start, end)
	s.Nil(err)
	s.Equal(startTime.String(), "2021-11-16 00:00:00 +0800 CST")
	s.Equal(endTime.String(), "2021-11-16 23:59:59.999999 +0800 CST")
}

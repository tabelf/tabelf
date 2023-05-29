package utils_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

var testSign = "whGDxFI6b0XxVYVRD+HWB+dti0h2afAp+lKuoYe4npVBEtE1YFVQ272iMZ8w2DDMytURstMi9nDMfW2PCrI6IB5jjCDvq1f7PP" +
	"/+3Q0IZBZLeQH5VvwQUL3lI0o0kGFjjlC+7hDTPPbvarj0QJLPxdW8ikCc4ZZAFbMxljshIvCMy2lLqqCjZgkQh8QSBSNJB8QBsuEhkQiC" +
	"hAI27rqqunns9cdIimvw2iClz/EF/Q7yCvjO8QfP6rcy1hVcgQsfxIZJ44EkMGfJAeKPCCmSPTuZYeb42CqGQZOpS8ahD7Vovr7DZ7zsxX" +
	"3w6fQ6UqwwkUoit4bfXaHvzHn3kDyoUw=="

type SignTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *SignTestSuite) Context() context.Context {
	return s.ctx
}

func (s *SignTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestSignTestSuite(t *testing.T) {
	suite.Run(t, new(SignTestSuite))
}

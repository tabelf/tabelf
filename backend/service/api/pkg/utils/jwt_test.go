package utils_test

import (
	"context"
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type JWTContextTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *JWTContextTestSuite) Context() context.Context {
	return s.ctx
}

func (s *JWTContextTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestJWTContextTestSuite(t *testing.T) {
	suite.Run(t, new(JWTContextTestSuite))
}

func (s *JWTContextTestSuite) TestGenerateToken() {
	result, err := utils.GetJwtToken("test", "zaihui", 10, utils.JWTClaims{OpenID: "test", AnonymousOpenID: "test"})
	s.Nil(err)
	s.NotEqual(result, "")
}

func (s *JWTContextTestSuite) TestIsValidateOfToken() {
	token, err := utils.GetJwtToken("test", "zaihui", 10, utils.JWTClaims{OpenID: "test", AnonymousOpenID: "test"})
	s.Nil(err)
	s.Run("test success", func() {
		tokenType, cal := utils.IsValidateOfToken(token, "test")
		s.Equal(utils.ValidateToken, tokenType)
		s.Equal(cal.OpenID, "test")
	})
	s.Run("test fail", func() {
		tokenType, _ := utils.IsValidateOfToken("test", "test")
		s.Equal(utils.BadToken, tokenType)
	})
	s.Run("test expire", func() {
		token2, err := utils.GetJwtToken("test", "zaihui", -1, utils.JWTClaims{OpenID: "test", AnonymousOpenID: "test"})
		s.Nil(err)
		tokenType, _ := utils.IsValidateOfToken(token2, "test")
		s.Equal(utils.ExpiredToken, tokenType)
	})
}

func (s *JWTContextTestSuite) TestParseToken() {
	token, err := utils.GetJwtToken("test", "zaihui", 10, utils.JWTClaims{OpenID: "test", AnonymousOpenID: "test"})
	s.Nil(err)
	result, err := utils.ParseToken(token, "test")
	s.Nil(err)
	s.Equal(result.OpenID, "test")
}

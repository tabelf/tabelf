package utils_test

import (
	"context"
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type AESContextTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *AESContextTestSuite) Context() context.Context {
	return s.ctx
}

func (s *AESContextTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestAESContextTestSuite(t *testing.T) {
	suite.Run(t, new(AESContextTestSuite))
}

func (s *AESContextTestSuite) TestPKCS7UnPadding() {
	data := []byte("WSuXADkLG3AjKXa6uf6B7oKxM8lr/SZmoJlOC0poCTfn4HtxKttb6iP" +
		"gArmUoOljTcZo06euX2vwP8aNl7DV2LxfbExwt+uUTd1PTRL/C3joheelo")
	res := utils.PKCS7UnPadding(data)
	s.Equal(res, []byte{87, 83})
}

package utils_test

import (
	"context"
	"strconv"
	"tabelf/backend/service/api/pkg/utils"
	"testing"

	"github.com/stretchr/testify/suite"
	"tabelf/backend/service/api/pkg/utils"
)

type ConvertorTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *ConvertorTestSuite) Context() context.Context {
	return s.ctx
}

func (s *ConvertorTestSuite) SetupTest() {
	s.ctx = context.Background()
}

func TestConvertorTestSuite(t *testing.T) {
	suite.Run(t, new(ConvertorTestSuite))
}

func (s *ConvertorTestSuite) TestToInt() {
	testSuccessData := []interface{}{
		"123", "-123", 123,
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float32(12.3), float64(12.3),
	}
	testFailData := []interface{}{
		"qwe", false, "666666666666666666666666666666666666666",
	}

	expectData := []int{123, -123, 123, 123, 123, 123, 123, 123, 12, 12}

	for i := 0; i < len(testSuccessData); i++ {
		actual, err := utils.ToInt(testSuccessData[i])
		s.NoError(err)
		s.Equal(expectData[i], actual)
	}

	expectError := []error{strconv.ErrSyntax, utils.ErrDataType, strconv.ErrRange}
	for i := 0; i < len(testFailData); i++ {
		actual, err := utils.ToInt(testFailData[i])
		s.ErrorIs(err, expectError[i])
		s.Equal(0, actual)
	}
}

func (s *ConvertorTestSuite) TestToString() {
	m := make(map[string]int)
	m["b"] = 1
	m["c"] = 2
	m["j"] = 1
	m["l"] = 1

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "testStruct"}

	testData := []interface{}{
		"", nil,
		int(0), int8(1), int16(-1), int32(123), int64(123),
		uint(123), uint8(123), uint16(123), uint32(123), uint64(123),
		float64(12.3), float32(12.3),
		true, false,
		[]int{
			1, 2, 3,
		},
		m, aStruct,
		[]byte{
			104, 101, 108, 108, 111,
		},
	}

	expectData := []string{
		"", "",
		"0", "1", "-1",
		"123", "123", "123", "123", "123", "123", "123",
		"12.3", "12.300000190734863",
		"true", "false",
		"[1,2,3]", "{\"b\":1,\"c\":2,\"j\":1,\"l\":1}", "{\"Name\":\"testStruct\"}", "hello",
	}

	for i := 0; i < len(testData); i++ {
		actual := utils.ToString(testData[i])
		s.Equal(expectData[i], actual)
	}
}

func (s *ConvertorTestSuite) TestToJson() {
	m := map[string]int{
		"b": 1, "c": 2, "j": 1, "l": 1,
	}
	mJSONStr, err := utils.ToJSON(m)
	s.NoError(err)
	s.Equal("{\"b\":1,\"c\":2,\"j\":1,\"l\":1}", mJSONStr)

	type TestStruct struct {
		Name string
	}
	aStruct := TestStruct{Name: "testStruct"}
	sJSONStr, err := utils.ToJSON(aStruct)
	s.NoError(err)
	s.Equal("{\"Name\":\"testStruct\"}", sJSONStr)
}

func (s *ConvertorTestSuite) TestStructToMap() {
	type TestStruct struct {
		Name string `json:"name"`
		age  int
	}
	testStruct := TestStruct{
		"test",
		100,
	}
	pm, err := utils.StructToMap(testStruct)
	s.NoError(err)
	expected := map[string]interface{}{
		"name": "test",
	}
	s.Equal(expected, pm)
}

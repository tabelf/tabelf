package testutils

import (
	"errors"

	"github.com/alicebob/miniredis/v2"
)

var ErrMock = errors.New("mock error")

func MockRedisServer() *miniredis.Miniredis {
	s := miniredis.NewMiniRedis()
	if err := s.Start(); err != nil {
		panic(err)
	}
	return s
}

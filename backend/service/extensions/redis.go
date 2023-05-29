package extensions

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
	"go.elastic.co/apm/module/apmgoredisv8"
)

type RedisConf struct {
	Enable bool   `json:"enable"`
	Addr   string `json:"addr"`
}

type RedisExt struct {
	RedisConf
	RedisClient *redis.Client
}

type SetOptions struct {
	Key   string
	Extra interface{}
}

func NewRedisExt(c RedisConf) *RedisExt {
	return &RedisExt{RedisConf: c, RedisClient: nil}
}

func (e *RedisExt) Init() error {
	if !e.Enable {
		return nil
	}
	ctx := context.Background()
	e.RedisClient = redis.NewClient(&redis.Options{Addr: e.Addr})
	e.RedisClient.AddHook(apmgoredisv8.NewHook())
	_, err := e.RedisClient.Ping(ctx).Result()
	return err
}

func (e *RedisExt) Close() error {
	return e.RedisClient.Close()
}

func (e *RedisExt) OpenAPIMiddleWare(next http.Handler) http.Handler {
	return next
}

func (e *RedisExt) Client(ctx context.Context) *redis.Client {
	return e.RedisClient.WithContext(ctx)
}

func (e *RedisExt) GetOrSet(
	ctx context.Context, opt *SetOptions, setAction func(*SetOptions) (string, error),
) (value interface{}, err error) {
	client := e.Client(ctx)
	result := client.Get(ctx, opt.Key)
	if result.Err() != nil {
		value, err = setAction(opt)
	} else {
		value = result.Val()
	}
	return
}

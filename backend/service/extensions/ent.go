package extensions

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jmoiron/sqlx"
)

type EntConf struct {
	Driver       string `json:"driver"`
	URL          string `json:"url"`
	MaxOpenConns int    `json:"max_open_conns"`
	MaxIdleConns int    `json:"max_idle_conns"`
}

type EntClient interface {
	Close() error
}

type EntExt struct {
	EntConf
	Client    EntClient
	NewClient func(*CacheDriver) EntClient
	drv       *CacheDriver
	redis     *RedisExt
	RowSQL    *sqlx.DB
}

func NewEntExt(c EntConf, newClient func(*CacheDriver) EntClient, redis *RedisExt) *EntExt {
	return &EntExt{EntConf: c, NewClient: newClient, redis: redis}
}

func (e *EntExt) Init() (err error) {
	var db *sql.DB
	db, err = sql.Open(e.EntConf.Driver, e.URL)
	if err != nil {
		return err
	}
	e.RowSQL = sqlx.NewDb(db, e.EntConf.Driver)
	db.SetMaxOpenConns(e.MaxOpenConns)
	db.SetMaxIdleConns(e.MaxIdleConns)
	drv := entsql.OpenDB(e.EntConf.Driver, db)
	e.drv = NewCacheDriver(
		drv,
		entcache.TTL(30*time.Second),
		entcache.Levels(
			entcache.NewLRU(256),
			entcache.NewRedis(e.redis.RedisClient),
		),
	)
	e.Client = e.NewClient(e.drv)
	return nil
}

func (e *EntExt) Close() error {
	return e.Client.Close()
}

func (e *EntExt) Driver() *entcache.Driver {
	return &e.drv.Driver
}

func (e *EntExt) OpenAPIMiddleWare(next http.Handler) http.Handler {
	return next
}

type CacheOpts struct {
	isCache bool
}

var cacheKey CacheOpts

type CacheDriver struct {
	entcache.Driver
}

func NewCacheDriver(drv dialect.Driver, opts ...entcache.Option) *CacheDriver {
	options := &entcache.Options{Hash: entcache.DefaultHash, Cache: entcache.NewLRU(0)}
	for _, opt := range opts {
		opt(options)
	}
	return &CacheDriver{entcache.Driver{Driver: drv, Options: options}}
}

func (d *CacheDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	if c, ok := ctx.Value(cacheKey).(*CacheOpts); ok {
		if c.isCache {
			return d.Driver.Query(ctx, query, args, v)
		}
	}
	return d.Driver.Driver.Query(ctx, query, args, v)
}

func Cache(ctx context.Context) context.Context {
	c, ok := ctx.Value(cacheKey).(*CacheOpts)
	if !ok {
		return context.WithValue(ctx, cacheKey, &CacheOpts{isCache: true})
	}
	c.isCache = true
	return ctx
}

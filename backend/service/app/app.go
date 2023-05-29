package app

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 本地用的MySql.
	"github.com/hashicorp/go-multierror"
	_ "github.com/mattn/go-sqlite3"           // ut用的sqlite.
	_ "go.elastic.co/apm/module/apmsql/mysql" // 测试、生产用的是 apm MySQL.
	schema "tabelf/backend/gen/entschema"
	"tabelf/backend/service/extensions"
)

var ErrAppInit = errors.New("app initialize error")

var (
	Basic     BasicConf
	Wechat    WechatConf
	Cos       CosConf
	Email     EmailConfig
	Log       *extensions.LogExt
	Redis     *extensions.RedisExt
	Ent       *extensions.EntExt
	EntClient *schema.Client
)

type Ext struct {
	Key       string
	Extension extensions.Extension
}

func getExtensions(conf Config, env string) []Ext {
	Basic = conf.Basic
	Wechat = conf.Wechat
	Cos = conf.Cos
	Email = conf.Email
	Log = extensions.NewLogExt(conf.LOG)
	Redis = extensions.NewRedisExt(conf.Redis)
	entNewClient := func(drv *extensions.CacheDriver) extensions.EntClient {
		return schema.NewClient(schema.Driver(drv), schema.Log(Log.UnionLogger.Debug))
	}
	Ent = extensions.NewEntExt(conf.DB, entNewClient, Redis)
	return []Ext{
		{Key: "log", Extension: Log},
		{Key: "redis", Extension: Redis},
		{Key: "entschema", Extension: Ent},
	}
}

func InitExtensions(conf Config, env string) error {
	var allErrs error
	for _, ext := range getExtensions(conf, env) {
		if err := ext.Extension.Init(); err != nil {
			allErrs = multierror.Append(allErrs, fmt.Errorf("ErrAppInit %w: %s", ErrAppInit, ext.Key), err)
		}
	}
	if allErrs == nil {
		var ok bool
		EntClient, ok = Ent.Client.(*schema.Client)
		if !ok {
			return fmt.Errorf("ErrAppInit %w: fail to initialize entschema client", ErrAppInit)
		}
	}
	return allErrs
}

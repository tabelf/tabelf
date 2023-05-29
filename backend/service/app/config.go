package app

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tabelf/backend/service/extensions"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

type WechatConf struct {
	AppID     string `json:"appid"`
	AppSecret string `json:"appsecret"`
	QrExpired int    `json:"qr_expired"`
}

type CosConf struct {
	SecretID  string `json:"secret_id"`
	SecretKey string `json:"secret_key"`
}

type EmailConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
}

type BasicConf struct {
	Domain string `json:"domain"`
	Env    string `json:"env"`
}

type Config struct {
	rest.RestConf
	LOG    extensions.LogConf   `json:"zlog"`
	Redis  extensions.RedisConf `json:"redis"`
	Jwt    extensions.JWTConf   `json:"jwt"`
	DB     extensions.EntConf   `json:"db"`
	Basic  BasicConf            `json:"basic"`
	Wechat WechatConf           `json:"wechat"`
	Cos    CosConf              `json:"cos"`
	Email  EmailConfig          `json:"email"`
}

type Configs struct {
	Default Config `json:"default"`
}

func LoadConfig(env string) Config {
	curDir, err := os.Getwd()
	if err != nil {
		log.Panicf("no current folder")
	}
	index := strings.Index(curDir, "tabelf")
	root := curDir[:index+len("tabelf")] + "/backend"
	var configFile string
	if flag.Lookup("f") == nil {
		flag.StringVar(&configFile, "f", filepath.Join(root, "config.yml"), "the config file")
	}
	var ok bool
	configFile, ok = flag.Lookup("f").Value.(flag.Getter).Get().(string)
	if !ok {
		log.Panicf("init config file")
	}
	flag.Parse()

	logx.SetLevel(logx.ErrorLevel)
	var c Configs
	conf.MustLoad(configFile, &c)
	return c.Default
}

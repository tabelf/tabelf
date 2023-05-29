package rpc

import (
	"context"
	"fmt"

	"github.com/zaihui/go-hutils"
	"github.com/zeromicro/go-zero/core/discov"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/prometheus"
	sev "github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"tabelf/backend/service/app"
	"tabelf/backend/service/rpc/service"
)

const (
	prometheusPort = 9001
)

func CreateServer(env string) (*zrpc.RpcServer, app.Config) {
	envConfig := app.LoadConfig(env)

	s := zrpc.MustNewServer(zrpc.RpcServerConf{
		Etcd: discov.EtcdConf{
			Hosts: envConfig.RPCServer.Etcd.Hosts,
			Key:   envConfig.RPCServer.Etcd.Key,
		},
		ServiceConf: sev.ServiceConf{
			Name: "binghuang.rpc",
			Log: logx.LogConf{
				Path: ".",
				Mode: "console",
			},
			Prometheus: prometheus.Config{
				Port: prometheusPort,
			},
		},
		ListenOn: envConfig.RPCServer.ListenOn,
		Timeout:  envConfig.RPCServer.Timeout,
	}, func(grpcServer *grpc.Server) {
		configureAPI(grpcServer, &GodServer{})
		reflection.Register(grpcServer)
	})
	defer s.Stop()

	err := app.InitExtensions(envConfig, env)
	if err != nil {
		app.Log.Panic(context.Background(), fmt.Errorf("init extensions failed: %w", err))
	}
	// 加载一下中间件
	s.AddUnaryInterceptors(app.APM.GrpcMiddleWare())
	s.AddUnaryInterceptors(app.Sentry.GrpcMiddleWare())
	s.AddUnaryInterceptors(hutils.NewUnaryServerAccessLogInterceptor(app.Log.UnionLogger))

	return s, envConfig
}

func StartGrpc(env string) {
	svr, cfg := CreateServer(env)
	fmt.Printf("Starting rpc server at %s...\n", cfg.RPCServer.ListenOn)
	app.Log.Track(context.Background(), fmt.Sprintf("Starting rpc server at %s...\n", cfg.RPCServer.ListenOn))
	svr.Start()
}

func configureAPI(s *grpc.Server, impls *GodServer) {
}

type GodServer struct {
	service.BaseService
}

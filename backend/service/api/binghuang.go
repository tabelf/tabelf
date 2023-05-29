package service

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"
	"unsafe"

	"github.com/RichardKnop/machinery/v2"
	"github.com/robfig/cron/v3"
	"github.com/zeromicro/go-zero/core/proc"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/go-zero/rest/router"

	"tabelf/backend/service/api/internal/handler"
	"tabelf/backend/service/api/internal/middleware"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/app"
	bcron "tabelf/backend/service/cron"
	"tabelf/backend/service/testutils"
)

func CreateServer(env string) (*rest.Server, *cron.Cron, *svc.ServiceContext) {
	envConfig := app.LoadConfig(env)
	ctx := svc.NewServiceContext(envConfig, env)
	server := rest.MustNewServer(envConfig.RestConf,
		rest.WithNotAllowedHandler(middleware.NewCorsMiddleware().Handler()))
	defer server.Stop()

	err := app.InitExtensions(envConfig, env)
	if err != nil {
		app.Log.Panic(context.Background(), fmt.Errorf("init extensions failed: %w", err))
	}
	// 加载一下中间件
	server.Use(middleware.NewCorsMiddleware().Handle)

	//server.Use(rest.ToMiddleware(app.Sentry.OpenAPIMiddleWare))
	server.Use(rest.ToMiddleware(app.Log.OpenAPIMiddleWare))

	handler.RegisterHandlers(server, ctx)
	proc.SetTimeToForceQuit(time.Duration(ctx.Config.RestConf.Timeout) * time.Millisecond)

	// 启动 cron
	cronjob := bcron.SetupCronJob(envConfig)

	return server, cronjob, ctx
}

func StartHTTP(env string) {
	server, cronjob, ctx := CreateServer(env)
	log.Printf("Starting server at %s:%d...", ctx.Config.Host, ctx.Config.Port)
	app.Log.Track(context.Background(), fmt.Sprintf("Starting server at %s:%d...", ctx.Config.Host, ctx.Config.Port))
	cronjob.Start()
	server.Start()
}

type TestApp struct {
	SVC    *svc.ServiceContext
	Router httpx.Router
}

func CreateTestApp() *TestApp {
	s := testutils.MockRedisServer()
	env := "testing"

	envConfig := app.LoadConfig(env)
	envConfig.Redis.Addr = s.Addr()

	ctx := svc.NewServiceContext(envConfig, env)
	server := rest.MustNewServer(envConfig.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)
	srvRouter := router.NewRouter()

	// UnsafePointer 是1.18的函数，升1.18后可替换
	g := (*Engine)(unsafe.Pointer(reflect.ValueOf(*server).Field(0).Pointer()))
	for _, r := range g.Routes {
		for _, route := range r.Routes {
			if err := srvRouter.Handle(route.Method, route.Path, route.Handler); err != nil {
				app.Log.Panic(context.Background(), fmt.Errorf("init test router failed: %w", err))
			}
		}
	}

	err := app.InitExtensions(envConfig, env)
	if err != nil {
		app.Log.Panic(context.Background(), fmt.Errorf("init extensions failed: %w", err))
	}
	// RegisterTasks(app.Machinery.Server)

	err = app.EntClient.Schema.Create(context.Background())
	if err != nil {
		app.Log.Panic(context.Background(), fmt.Errorf("create test app failed: %w", err))
	}
	return &TestApp{
		SVC:    ctx,
		Router: srvRouter,
	}
}

// RegisterTasks 注册异步任务.
func RegisterTasks(server *machinery.Server) {
	err := server.RegisterTasks(map[string]interface{}{
		// "RefundOrder":                  tasks.RefundOrder,
		// "SyncVideoID":                  models.SyncVideoID,
	})
	if err != nil {
		app.Log.Fatal(context.Background(), err)
	}
}

// SignatureSetting go-zero的struct.
// 以下是为了test，映射go-zero的route.
// 如果go-zero的结构变动，这里可能需要修改.
type SignatureSetting struct {
	rest.SignatureConf
	Enabled bool
}

type JwtSetting struct {
	Enabled    bool
	Secret     string
	PrevSecret string
}

type FeaturedRoutes struct {
	Timeout   time.Duration
	Priority  bool
	Jwt       JwtSetting
	Signature SignatureSetting
	Routes    []rest.Route
	MaxBytes  int64
}

type Engine struct {
	Conf   rest.RestConf
	Routes []FeaturedRoutes
}

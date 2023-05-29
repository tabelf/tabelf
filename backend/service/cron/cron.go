package cron

import (
	"context"
	"fmt"
	"time"

	"github.com/1278651995/cron_ui"
	"github.com/robfig/cron/v3"
	"tabelf/backend/service/app"
	"tabelf/backend/service/cron/tasks"
)

type JobFunc func(ctx tasks.JobContext, config app.Config)

type Job struct {
	Spec string
	Func JobFunc
	Name string
}

var EntriesInfo = map[cron.EntryID]cron_ui.JobInfo{}

// 需要执行的定时任务.
var jobs = []Job{
	// {Spec: "* * * * *", Func: tasks.SayHello, Name: "HelloWorld"},
	// Minute | Hour | Dom | Month | Dow
	{Spec: "*/30 * * * *", Func: tasks.OrderTimeoutCancel, Name: "OrderTimeoutCancel"},
	// 凌晨1点, 全体
	{Spec: "0 1 * * *", Func: tasks.SendNoticeMessage, Name: "SendNoticeMessage"},
	// 凌晨1点
	{Spec: "0 1 * * *", Func: tasks.CommunityUrlRecord, Name: "CommunityUrlRecord"},
	// 凌晨2点
	{Spec: "0 2 * * *", Func: tasks.AutoCreateStation, Name: "AutoCreateStation"},
	// 凌晨5点
	{Spec: "0 5 * * *", Func: tasks.BaiDuUrlCrawl, Name: "BaiDuUrlCrawl"},
	// 凌晨4点
	{Spec: "0 4 * * *", Func: tasks.UpdateGoodArticleHot, Name: "UpdateGoodArticleHot"},
}

func cronjobWithApmTrace(f func(ctx tasks.JobContext, config app.Config), config app.Config) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()
	f(tasks.NewJobContext(ctx), config)
}

// SetupCronJob 初始化cronjob.
func SetupCronJob(config app.Config) *cron.Cron {
	c := cron.New(cron.WithChain(cron.Recover(NewLogger())))
	for _, job := range jobs {
		f := job.Func
		entryID, err := c.AddFunc(job.Spec, func() { cronjobWithApmTrace(f, config) })
		if err != nil {
			app.Log.Fatal(context.TODO(), err)
		}
		EntriesInfo[entryID] = cron_ui.JobInfo{
			Name: job.Name,
			Spec: job.Spec,
		}
	}
	return c
}

// Logger 自定义Cron.Logger.
type Logger struct {
	ctx context.Context
}

func NewLogger() Logger {
	return Logger{ctx: context.Background()}
}

// Info level为debug.
func (l Logger) Info(msg string, keysAndValues ...interface{}) {
	app.Log.Track(l.ctx, fmt.Sprintf("%s %+v", msg, keysAndValues))
}

// Error level为error.
func (l Logger) Error(err error, msg string, keysAndValues ...interface{}) {
	app.Log.Error(l.ctx, fmt.Errorf("%w %s %+v", err, msg, keysAndValues))
}

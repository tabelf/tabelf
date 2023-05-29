package tasks

import (
	"log"
	entaccount "tabelf/backend/gen/entschema/account"
	entnotice "tabelf/backend/gen/entschema/notice"
	"tabelf/backend/service/api/models"
	"tabelf/backend/service/app"
)

// SendNoticeMessage 全局消息通知公告功能.
func SendNoticeMessage(jobCtx JobContext, config app.Config) {
	log.Print("send global message notice start.")
	ctx := jobCtx.Context
	notices, err := app.EntClient.Notice.Query().Where(
		entnotice.Process(false),
		entnotice.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return
	}
	for _, notice := range notices {
		accounts, err := app.EntClient.Account.Query().Where(
			entaccount.DeactivatedAtIsNil(),
		).All(ctx)
		if err != nil {
			return
		}
		for _, account := range accounts {
			// 已系统的名义发给每个人
			if err = models.SendMessage(
				ctx,
				app.SystemMessage,
				app.SystemAccount,
				account.UID,
				notice.Content,
			); err != nil {
				app.Log.Error(ctx, err)
			}
		}
		if err = app.EntClient.Notice.Update().SetProcess(true).Where(
			entnotice.ID(notice.ID),
		).Exec(ctx); err != nil {
			return
		}
	}
}

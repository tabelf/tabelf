package feedback

import (
	"context"
	"encoding/json"
	"github.com/jordan-wright/email"
	"net/smtp"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateFeedbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateFeedbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateFeedbackLogic {
	return &CreateFeedbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateFeedbackLogic) CreateFeedback(req *types.CreateFeedbackRequest) (resp *types.CreateFeedbackResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	if !app.StringSliceContains([]string{
		app.OrderFeedbackCategory,
		app.BugFeedbackCategory,
		app.OptimizationFeedbackCategory,
		app.OtherFeedbackCategory,
	}, req.Category) {
		return nil, app.ErrPersonalFeedbackCategoryNotExist(l.ctx)
	}
	if req.Category == app.OrderFeedbackCategory {
		feedback, err := app.EntClient.PayOrderFeedback.Create().
			SetOrderNumber(req.OrderNumber).
			SetCategory(req.Category).
			SetDescription(req.Description).
			SetUserUID(req.UserUID).
			Save(l.ctx)
		body, err := json.Marshal(feedback)
		if err != nil {
			return nil, err
		}
		e := &email.Email{
			To:      []string{app.OrderFeedbackSystemUsername},
			From:    app.Email.Username,
			Subject: app.EmailOrderFeedbackSubject,
			Text:    body,
		}
		if err = e.Send(app.Email.Addr+":"+app.EmailSSLPorts, smtp.PlainAuth(
			"",
			app.Email.Username,
			app.Email.Password,
			app.Email.Addr,
		)); err != nil {
			return nil, err
		}
		return &types.CreateFeedbackResponse{
			Message: app.HttpOK,
		}, nil
	}
	if err := app.EntClient.Feedback.Create().
		SetCategory(req.Category).
		SetDescription(req.Description).
		SetUserUID(req.UserUID).
		Exec(l.ctx); err != nil {
		return nil, err
	}
	return &types.CreateFeedbackResponse{
		Message: app.HttpOK,
	}, nil
}

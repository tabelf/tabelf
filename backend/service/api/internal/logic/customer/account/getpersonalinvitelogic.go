package account

import (
	"context"
	"fmt"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entinvite "tabelf/backend/gen/entschema/invite"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPersonalInviteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalInviteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalInviteLogic {
	return &GetPersonalInviteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalInviteLogic) GetPersonalInvite(req *types.GetPersonalInviteRequest) (resp *types.GetPersonalInviteResponse, err error) {
	invites, err := app.EntClient.Invite.Query().Where(
		entinvite.ReferralUID(req.UserUID),
		entinvite.DeactivatedAtIsNil(),
	).Order(entschema.Desc(entinvite.FieldCreatedAt)).All(l.ctx)
	if err != nil {
		return nil, err
	}
	maxNum := 20
	inviteUserUID := make([]string, 0)
	for _, invite := range invites {
		if maxNum > 0 {
			inviteUserUID = append(inviteUserUID, invite.RefereeUID)
			maxNum--
		}
	}
	achievements := make([]*types.Achievement, 0)
	if len(inviteUserUID) != 0 {
		accounts, err := app.EntClient.Account.Query().Where(
			entaccount.UIDIn(inviteUserUID...),
			entaccount.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		accountMap := make(map[string]*entschema.Account)
		for _, account := range accounts {
			accountMap[account.UID] = account
		}
		for _, uid := range inviteUserUID {
			if account, ok := accountMap[uid]; ok {
				achievements = append(achievements, &types.Achievement{
					UserUID:   account.UID,
					Image:     account.Image,
					UserName:  account.Nickname,
					CreatedAt: app.GetTime(account.CreatedAt),
				})
			}
		}
	}
	return &types.GetPersonalInviteResponse{
		InviteURL:    fmt.Sprintf("%s/r/%s", app.Basic.Domain, req.UserUID),
		Earned:       len(invites) * app.InviteRewardFileNum,
		Invited:      len(invites),
		Achievements: achievements,
	}, nil
}

package message

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entmessage "tabelf/backend/gen/entschema/message"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessagesLogic {
	return &GetMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessagesLogic) GetMessages(req *types.GetMessagesRequest) (resp *types.GetMessagesResponse, err error) {
	messages, err := app.EntClient.Message.Query().Where(
		entmessage.UserUID(req.UserUID),
		entmessage.DeactivatedAtIsNil(),
		func(s *sql.Selector) {
			if req.MsgType == 0 {
				s.Where(sql.EQ(entmessage.FieldHasRead, req.MsgType))
			}
		},
	).Order(entschema.Desc(entmessage.FieldCreatedAt)).
		All(l.ctx)
	if err != nil {
		return nil, err
	}
	userUIDs := make([]string, 0)
	for _, msg := range messages {
		userUIDs = append(userUIDs, msg.PromoterUID)
	}
	accounts, err := app.EntClient.Account.Query().Where(
		entaccount.UIDIn(userUIDs...),
		entaccount.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*entschema.Account, 0)
	for _, account := range accounts {
		userMap[account.UID] = account
	}
	userMessages := make([]*types.UserMessage, 0)
	unread := 0
	for _, msg := range messages {
		if user, ok := userMap[msg.PromoterUID]; ok {
			userMessages = append(userMessages, &types.UserMessage{
				UID:           msg.UID,
				PromoterName:  user.Nickname,
				PromoterImage: user.Image,
				Description:   msg.Description,
				HasRead:       msg.HasRead,
				CreatedAt:     app.GetTime(msg.CreatedAt),
			})
			if !msg.HasRead {
				unread++
			}
		}
	}
	return &types.GetMessagesResponse{
		Unread:       unread,
		UserMessages: userMessages,
	}, nil
}

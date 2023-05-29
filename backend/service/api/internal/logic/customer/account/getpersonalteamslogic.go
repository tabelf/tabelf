package account

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	entteam "tabelf/backend/gen/entschema/team"
	entteamgroup "tabelf/backend/gen/entschema/teamgroup"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetPersonalTeamsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalTeamsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalTeamsLogic {
	return &GetPersonalTeamsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalTeamsLogic) GetPersonalTeams(req *types.GetPersonalTeamsRequest) (resp *types.GetPersonalTeamsResponse, err error) {
	groups, err := app.EntClient.TeamGroup.Query().Where(
		entteamgroup.UserUID(req.UserUID),
		entteamgroup.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	personalTeams := make([]*types.PersonalTeam, 0)
	if len(groups) != 0 {
		groupUIDs := make([]string, 0)
		for _, group := range groups {
			groupUIDs = append(groupUIDs, group.TeamUID)
		}
		teams, err := app.EntClient.Team.Query().Where(
			entteam.UIDIn(groupUIDs...),
			entteam.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		for _, team := range teams {
			personalTeams = append(personalTeams, &types.PersonalTeam{
				TeamUID:   team.UID,
				TeamName:  team.Name,
				ExpiredAt: app.GetTime(team.ExpiredAt),
			})
		}
	}
	return &types.GetPersonalTeamsResponse{
		PersonalTeams: personalTeams,
	}, nil
}

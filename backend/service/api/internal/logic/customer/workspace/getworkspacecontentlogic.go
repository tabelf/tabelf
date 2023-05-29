package workspace

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/service/api/internal/logic/base"

	"tabelf/backend/gen/entschema"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetWorkspaceContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetWorkspaceContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWorkspaceContentLogic {
	return &GetWorkspaceContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetWorkspaceContentLogic) GetWorkspaceContent(req *types.GetWorkspaceContentRequest) (resp *types.GetWorkspaceContentResponse, err error) {
	folder, shareLink, personalWorkspaces, activeWorkspaceUIDs, err := base.WorkspaceContent(l.ctx, req.FolderUID)
	if err != nil {
		return nil, err
	}
	authority := app.ShareEditAuthority
	if folder.UserUID != req.UserUID {
		collaboration, err := app.EntClient.Collaboration.Query().Where(
			entcollaboration.UserUID(req.UserUID),
			entcollaboration.FolderUID(req.FolderUID),
			entcollaboration.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return nil, app.ErrPersonalFolderCollNotExist(l.ctx)
			}
			return nil, err
		}
		authority = collaboration.Authority
	}
	return &types.GetWorkspaceContentResponse{
		Authority:           authority,
		ShareUID:            shareLink.UID,
		FolderUID:           folder.UID,
		FolderName:          folder.FolderName,
		FolderNumber:        folder.FolderNumber,
		ActiveWorkspaceUIDs: activeWorkspaceUIDs,
		PersonalWorkspaces:  personalWorkspaces,
	}, nil
}

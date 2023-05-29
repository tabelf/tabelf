package folder

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetPersonalFoldersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPersonalFoldersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPersonalFoldersLogic {
	return &GetPersonalFoldersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPersonalFoldersLogic) GetPersonalFolders(req *types.GetPersonalFoldersRequest) (resp *types.GetPersonalFoldersResponse, err error) {
	folders, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UserUID(req.UserUID),
		entpersonalfolder.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	personalFolders := make([]*types.PersonalFolder, len(folders))
	for i, folder := range folders {
		personalFolders[i] = &types.PersonalFolder{
			FolderUID:    folder.UID,
			FolderNumber: folder.FolderNumber,
			FolderName:   folder.FolderName,
		}
	}
	return &types.GetPersonalFoldersResponse{
		Folders: personalFolders,
	}, nil
}

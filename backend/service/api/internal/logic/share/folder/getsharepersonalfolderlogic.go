package folder

import (
	"context"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSharePersonalFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSharePersonalFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSharePersonalFolderLogic {
	return &GetSharePersonalFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSharePersonalFolderLogic) GetSharePersonalFolder(req *types.GetSharePersonalFolderRequest) (resp *types.GetSharePersonalFolderResponse, err error) {
	// 与他人协作
	followers, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.UserUID(req.UserUID),
		entcollaboration.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	folderUIDs := make([]string, 0)
	shareFolderMap := map[string]string{}
	authorityMap := map[string]string{}
	for _, follower := range followers {
		folderUIDs = append(folderUIDs, follower.FolderUID)
		authorityMap[follower.FolderUID] = follower.Authority
		shareFolderMap[follower.FolderUID] = follower.ShardUID
	}
	if len(folderUIDs) != 0 {
		folders, err := app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UIDIn(folderUIDs...),
			entpersonalfolder.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		personalFolders := make([]*types.ShareFolder, len(folders))
		for i, folder := range folders {
			authority := "0"
			if _, ok := authorityMap[folder.UID]; ok {
				authority = authorityMap[folder.UID]
			}
			personalFolders[i] = &types.ShareFolder{
				ShareUID:     shareFolderMap[folder.UID],
				HasOwner:     folder.UserUID == req.UserUID,
				Authority:    authority,
				FolderUID:    folder.UID,
				FolderNumber: folder.FolderNumber,
				FolderName:   folder.FolderName,
			}
		}
		return &types.GetSharePersonalFolderResponse{
			Folders: personalFolders,
		}, nil
	}
	return &types.GetSharePersonalFolderResponse{
		Folders: make([]*types.ShareFolder, 0),
	}, nil
}

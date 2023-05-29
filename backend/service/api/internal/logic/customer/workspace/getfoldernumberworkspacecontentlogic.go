package workspace

import (
	"context"
	"sort"
	"tabelf/backend/gen/entschema"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFolderNumberWorkspaceContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFolderNumberWorkspaceContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFolderNumberWorkspaceContentLogic {
	return &GetFolderNumberWorkspaceContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFolderNumberWorkspaceContentLogic) GetFolderNumberWorkspaceContent(req *types.GetFolderNumberWorkspaceContentRequest) (resp *types.GetFolderNumberWorkspaceContentResponse, err error) {
	if app.IsBlank(req.FolderNumber) {
		return nil, app.ErrPersonalFolderContentNotExist(l.ctx)
	}
	folder, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UserUID(req.UserUID),
		entpersonalfolder.FolderNumber(req.FolderNumber),
		entpersonalfolder.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderNotExist(l.ctx)
		}
		return nil, err
	}
	workspaces, err := app.EntClient.Workspace.Query().Where(
		entworkspace.PersonalFolderUID(folder.UID),
		entworkspace.Type(app.PersonalWorkspaceType),
		entworkspace.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	workspaceUIDs := make([]string, 0)
	for _, workspace := range workspaces {
		workspaceUIDs = append(workspaceUIDs, workspace.UID)
	}
	links, err := app.EntClient.WebLink.Query().Where(
		entweblink.WorkspaceUIDIn(workspaceUIDs...),
		entweblink.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	sort.Slice(links, func(i, j int) bool {
		return links[i].Sequence < links[j].Sequence
	})
	workspaceLinkMap := make(map[string][]*types.PersonalWebLink, 0)
	for _, link := range links {
		webLink := &types.PersonalWebLink{
			LinkUID:     link.UID,
			Title:       link.Title,
			Image:       link.Image,
			Link:        link.Link,
			Description: link.Description,
			FileType:    link.FileType,
			Sequence:    link.Sequence,
		}
		if maps, ok := workspaceLinkMap[link.WorkspaceUID]; ok {
			workspaceLinkMap[link.WorkspaceUID] = append(maps, webLink)
		} else {
			workspaceLinkMap[link.WorkspaceUID] = []*types.PersonalWebLink{webLink}
		}
	}
	personalWorkspaces := make([]*types.PersonalWorkspace, 0)
	activeWorkspaceUIDs := make([]string, 0)
	for _, workspace := range workspaces {
		if workspace.IsOpen {
			activeWorkspaceUIDs = append(activeWorkspaceUIDs, workspace.UID)
		}
		if webLinks, ok := workspaceLinkMap[workspace.UID]; ok {
			personalWorkspaces = append(personalWorkspaces, &types.PersonalWorkspace{
				IsOpen:        workspace.IsOpen,
				WorkspaceUID:  workspace.UID,
				WorkspaceName: workspace.Name,
				WebLinks:      webLinks,
			})
		} else {
			personalWorkspaces = append(personalWorkspaces, &types.PersonalWorkspace{
				IsOpen:        workspace.IsOpen,
				WorkspaceUID:  workspace.UID,
				WorkspaceName: workspace.Name,
				WebLinks:      make([]*types.PersonalWebLink, 0),
			})
		}
	}
	return &types.GetFolderNumberWorkspaceContentResponse{
		FolderUID:           folder.UID,
		FolderName:          folder.FolderName,
		FolderNumber:        folder.FolderNumber,
		ActiveWorkspaceUIDs: activeWorkspaceUIDs,
		PersonalWorkspaces:  personalWorkspaces,
	}, nil
}

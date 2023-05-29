package workspace

import (
	"context"
	"sort"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/sync/errgroup"

	"tabelf/backend/gen/entschema"
	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type GetFolderNumberShareWorkspaceContentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFolderNumberShareWorkspaceContentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFolderNumberShareWorkspaceContentLogic {
	return &GetFolderNumberShareWorkspaceContentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFolderNumberShareWorkspaceContentLogic) GetFolderNumberShareWorkspaceContent(req *types.GetFolderNumberShareWorkspaceContentRequest) (resp *types.GetFolderNumberShareWorkspaceContentResponse, err error) {
	if app.IsBlank(req.FolderNumber) {
		return nil, app.ErrPersonalFolderContentNotExist(l.ctx)
	}
	collaboration, err := app.EntClient.Collaboration.Query().Where(
		entcollaboration.UserUID(req.UserUID),
		entcollaboration.FolderNumber(req.FolderNumber),
		entcollaboration.DeactivatedAtIsNil(),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrPersonalFolderCollNotExist(l.ctx)
		}
		return nil, err
	}
	var (
		g          errgroup.Group
		folder     *entschema.PersonalFolder
		workspaces []*entschema.Workspace
	)
	g.Go(func() (err error) {
		folder, err = app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.UID(collaboration.FolderUID),
			entpersonalfolder.FolderNumber(req.FolderNumber),
			entpersonalfolder.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrPersonalFolderNotExist(l.ctx)
			}
			return err
		}
		return nil
	})
	g.Go(func() (err error) {
		workspaces, err = app.EntClient.Workspace.Query().Where(
			entworkspace.PersonalFolderUID(collaboration.FolderUID),
			entworkspace.Type(app.PersonalWorkspaceType),
			entworkspace.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err = g.Wait(); err != nil {
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
			FileType:    link.FileType,
			Description: link.Description,
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
	return &types.GetFolderNumberShareWorkspaceContentResponse{
		Authority:           collaboration.Authority,
		ShareUID:            collaboration.ShardUID,
		FolderUID:           folder.UID,
		FolderName:          folder.FolderName,
		FolderNumber:        folder.FolderNumber,
		ActiveWorkspaceUIDs: activeWorkspaceUIDs,
		PersonalWorkspaces:  personalWorkspaces,
	}, nil
}

package account

import (
	"context"
	"tabelf/backend/gen/entschema"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"

	"entgo.io/ent/dialect/sql"
	"github.com/zeromicro/go-zero/core/logx"

	entcollaboration "tabelf/backend/gen/entschema/collaboration"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
)

type SearchWebLinksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchWebLinksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchWebLinksLogic {
	return &SearchWebLinksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchWebLinksLogic) SearchWebLinks(req *types.SearchWebLinksRequest) (resp *types.SearchWebLinksResponse, err error) {
	webLinks := make([]*types.SearchWebLink, 0)
	if app.IsBlank(req.Keyword) {
		return &types.SearchWebLinksResponse{
			SearchWebLinks: webLinks,
		}, nil
	}
	folders := make([]string, 0)
	if req.Type == 0 || req.Type == 1 {
		links, err := app.EntClient.WebLink.Query().Where(
			entweblink.UserUID(req.UserUID),
			entweblink.Or(
				func(s *sql.Selector) {
					s.Where(sql.Like(entweblink.FieldDescription, "%"+req.Keyword+"%"))
				},
				func(s *sql.Selector) {
					s.Where(sql.Like(entweblink.FieldTitle, "%"+req.Keyword+"%"))
				},
			),
			entweblink.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		for _, link := range links {
			folders = append(folders, link.FolderUID)
		}
		for _, link := range links {
			webLinks = append(webLinks, &types.SearchWebLink{
				FolderUID:   link.FolderUID,
				LinkUID:     link.UID,
				Title:       link.Title,
				Image:       link.Image,
				Link:        link.Link,
				Description: link.Description,
			})
		}
	}
	if req.Type == 0 || req.Type == 2 {
		// 查询我的协作文件
		followers, err := app.EntClient.Collaboration.Query().Where(
			entcollaboration.UserUID(req.UserUID),
			entcollaboration.DeactivatedAtIsNil(),
		).All(l.ctx)
		if err != nil {
			return nil, err
		}
		folderUIDs := make([]string, 0)
		for _, c := range followers {
			folderUIDs = append(folderUIDs, c.FolderUID)
		}
		if len(folderUIDs) != 0 {
			workspaces, err := app.EntClient.Workspace.Query().Where(
				entworkspace.PersonalFolderUIDIn(folderUIDs...),
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
			if len(workspaceUIDs) != 0 {
				links, err := app.EntClient.WebLink.Query().Where(
					entweblink.WorkspaceUIDIn(workspaceUIDs...),
					entweblink.Or(
						func(s *sql.Selector) {
							s.Where(sql.Like(entweblink.FieldDescription, "%"+req.Keyword+"%"))
						},
						func(s *sql.Selector) {
							s.Where(sql.Like(entweblink.FieldTitle, "%"+req.Keyword+"%"))
						},
					),
					entweblink.DeactivatedAtIsNil(),
				).All(l.ctx)
				if err != nil {
					return nil, err
				}
				for _, link := range links {
					folders = append(folders, link.FolderUID)
				}
				for _, link := range links {
					webLinks = append(webLinks, &types.SearchWebLink{
						FolderUID:   link.FolderUID,
						LinkUID:     link.UID,
						Title:       link.Title,
						Image:       link.Image,
						Link:        link.Link,
						Description: link.Description,
					})
				}
			}
		}
	}
	personalFolders, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UIDIn(folders...),
		entpersonalfolder.DeactivatedAtIsNil(),
	).All(l.ctx)
	if err != nil {
		return nil, err
	}
	folderMap := make(map[string]*entschema.PersonalFolder)
	for _, folder := range personalFolders {
		folderMap[folder.UID] = folder
	}
	for i, web := range webLinks {
		if f, ok := folderMap[web.FolderUID]; ok {
			webLinks[i].FolderNumber = f.FolderNumber
		}
	}
	return &types.SearchWebLinksResponse{
		SearchWebLinks: webLinks,
	}, nil
}

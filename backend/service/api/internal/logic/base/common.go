package base

import (
	"context"
	"github.com/badoux/goscraper"
	"net/url"
	"sort"
	"strings"
	"tabelf/backend/common"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entadmin "tabelf/backend/gen/entschema/admin"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	entsharelink "tabelf/backend/gen/entschema/sharelink"
	entweblink "tabelf/backend/gen/entschema/weblink"
	entworkspace "tabelf/backend/gen/entschema/workspace"
	"tabelf/backend/service/api/internal/middleware"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/app"
	"time"
)

func ValidateUserAuthority(ctx context.Context, userID string) (err error) {
	element, err := middleware.ContextElement(ctx)
	if err != nil || element.UID != userID {
		return app.ErrAccountOutOfAuthority(ctx)
	}
	return nil
}

// MembershipValidity 验证会员有效期. true 有效，false 无效
func MembershipValidity(account *entschema.Account) bool {
	if account == nil {
		return false
	}
	if app.StringSliceContains([]string{
		app.MonthMemberUser,
		app.YearMemberUser,
	}, account.MemberType) && time.Now().Before(account.MemberExpired) {
		return true
	}
	return false
}

// HasAdminAuthority 是否有管理权限.
func HasAdminAuthority(ctx context.Context, userUID string) (bool, error) {
	user, err := app.EntClient.Account.Query().Where(
		entaccount.UID(userUID),
		entaccount.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return false, app.ErrAccountInvalid(ctx)
		}
		return false, err
	}
	if !user.HasAdmin {
		return false, nil
	}

	return app.EntClient.Admin.Query().Where(
		entadmin.UserUID(userUID),
		entadmin.DeactivatedAtIsNil(),
	).Exist(ctx)
}

type WebLinkInfo struct {
	Host        string
	Title       string
	Description string
	Image       string
}

func GetWebLinkInfo(ctx context.Context, URL string) (info *WebLinkInfo, err error) {
	u, err := url.Parse(URL)
	if err != nil {
		return nil, app.ErrCustomerWebLinkInvalid(ctx)
	}
	s, err := goscraper.Scrape(URL, 5)
	if err != nil {
		return &WebLinkInfo{
			Host:        URL,
			Title:       u.Host,
			Description: "",
			Image:       "",
		}, nil
	}
	icon := s.Preview.Icon
	if !strings.HasPrefix(icon, "http") {
		if strings.HasPrefix(icon, "//") {
			icon = icon[2:]
		} else if strings.HasPrefix(icon, "/") {
			if len(s.Preview.Images) > 0 {
				icon = s.Preview.Images[0]
			} else {
				icon = u.Host + icon
			}
		}
	}
	if !strings.HasPrefix(icon, "http") {
		icon = app.HTTPProtoType + icon
	}
	return &WebLinkInfo{
		Host:        u.Host,
		Title:       s.Preview.Title,
		Description: s.Preview.Description,
		Image:       icon,
	}, nil
}

func CopyFolderContent(ctx context.Context, account *entschema.Account, folder *entschema.PersonalFolder) (folderNumber string, err error) {
	workspaces, err := app.EntClient.Workspace.Query().Where(
		entworkspace.PersonalFolderUID(folder.UID),
		entworkspace.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return "", err
	}

	for {
		folderNumber = app.RandomString(app.FolderNumberCount)
		exist, err := app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.FolderNumber(folderNumber),
			entpersonalfolder.UserUID(account.UID),
			entpersonalfolder.DeactivatedAtIsNil(),
		).Exist(ctx)
		if err != nil {
			return "", err
		}
		if !exist {
			break
		}
	}

	copyLinkCount := 0
	newFolderUID := common.NewUUID()
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		folderCreate := tx.PersonalFolder.Create().
			SetUID(newFolderUID).
			SetUserUID(account.UID).
			SetFolderName(folder.FolderName).
			SetFolderNumber(folderNumber)

		shareLinkCreate := tx.ShareLink.Create().
			SetFolderUID(newFolderUID).
			SetUserUID(account.UID).
			SetAuthority(app.ShareReadAuthority).
			SetValidDay(app.ForEverDay).
			SetExpiredAt(app.ParseTime(app.ForEverValid)).
			SetRecentAt(time.Now()).
			SetFolderNumber(folderNumber)

		workspaceBulks := make([]*entschema.WorkspaceCreate, 0)
		webLinkBulks := make([]*entschema.WebLinkCreate, 0)
		for _, workspace := range workspaces {
			newWorkspaceUID := common.NewUUID()
			workspaceCreate := tx.Workspace.Create().
				SetUID(newWorkspaceUID).
				SetName(workspace.Name).
				SetType(workspace.Type).
				SetUserUID(account.UID).
				SetPersonalFolderUID(newFolderUID).
				SetIsOpen(workspace.IsOpen)
			workspaceBulks = append(workspaceBulks, workspaceCreate)

			links, err := app.EntClient.WebLink.Query().Where(
				entweblink.WorkspaceUID(workspace.UID),
				entweblink.DeactivatedAtIsNil(),
			).All(ctx)
			if err != nil {
				return err
			}
			for _, link := range links {
				linkCreate := tx.WebLink.Create().
					SetTitle(link.Title).
					SetDescription(link.Description).
					SetImage(link.Image).
					SetLink(link.Link).
					SetSequence(link.Sequence).
					SetUserUID(account.UID).
					SetFileType(link.FileType).
					SetWorkspaceUID(newWorkspaceUID).
					SetFolderUID(newFolderUID)

				webLinkBulks = append(webLinkBulks, linkCreate)
				copyLinkCount++
			}
		}
		// 不为会员 或者 会员过期
		//if !app.StringSliceContains([]string{
		//	app.MonthMemberUser,
		//	app.YearMemberUser,
		//}, account.MemberType) || time.Now().After(account.MemberExpired) {
		//	if account.URLCount+copyLinkCount > account.URLLimit {
		//		return app.ErrCustomerWebLinkLimit(ctx)
		//	}
		//}
		if err = folderCreate.Exec(ctx); err != nil {
			return err
		}
		if err = shareLinkCreate.Exec(ctx); err != nil {
			return err
		}
		if err = tx.Workspace.CreateBulk(workspaceBulks...).
			Exec(ctx); err != nil {
			return err
		}
		if err = tx.WebLink.CreateBulk(webLinkBulks...).
			Exec(ctx); err != nil {
			return err
		}
		if err = tx.Account.Update().AddURLCount(copyLinkCount).
			Where(entaccount.ID(account.ID)).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", err
	}
	return folderNumber, nil
}

func WorkspaceContent(ctx context.Context, folderUID string) (*entschema.PersonalFolder, *entschema.ShareLink, []*types.PersonalWorkspace, []string, error) {
	folder, err := app.EntClient.PersonalFolder.Query().Where(
		entpersonalfolder.UID(folderUID),
		entpersonalfolder.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, nil, nil, nil, app.ErrPersonalFolderNotExist(ctx)
		}
		return nil, nil, nil, nil, err
	}
	shareLink, err := app.EntClient.ShareLink.Query().Where(
		entsharelink.FolderUID(folderUID),
		entsharelink.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	workspaces, err := app.EntClient.Workspace.Query().Where(
		entworkspace.PersonalFolderUID(folderUID),
		entworkspace.Type(app.PersonalWorkspaceType),
		entworkspace.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	workspaceUIDs := make([]string, 0)
	for _, workspace := range workspaces {
		workspaceUIDs = append(workspaceUIDs, workspace.UID)
	}
	links, err := app.EntClient.WebLink.Query().Where(
		entweblink.WorkspaceUIDIn(workspaceUIDs...),
		entweblink.DeactivatedAtIsNil(),
	).All(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
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
	return folder, shareLink, personalWorkspaces, activeWorkspaceUIDs, nil
}

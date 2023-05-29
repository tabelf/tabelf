package base

import (
	"context"
	"golang.org/x/sync/errgroup"
	"strings"
	entworkspace "tabelf/backend/gen/entschema/workspace"

	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/app"
)

func AddWebLink(ctx context.Context, url, title, userUID, workspaceUID string) (err error) {
	var (
		g errgroup.Group
		//account   *entschema.Account
		workspace *entschema.Workspace
		info      *WebLinkInfo
		count     int
	)
	g.Go(func() (err error) {
		_, err = app.EntClient.Account.Query().Where(
			entaccount.UID(userUID),
			entaccount.DeactivatedAtIsNil(),
		).First(ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrAccountInvalid(ctx)
			}
			return err
		}
		// 会员不做数量校验
		//if MembershipValidity(account) {
		//	return nil
		//}
		//if account.URLCount+1 > account.URLLimit {
		//	return app.ErrCustomerWebLinkLimit(ctx)
		//}
		return nil
	})
	g.Go(func() (err error) {
		if !strings.HasPrefix(url, "http") {
			url = app.HTTPProtoType + url
		}
		info, err = GetWebLinkInfo(ctx, url)
		return err
	})
	g.Go(func() (err error) {
		count, err = app.EntClient.WebLink.Query().Where(
			entweblink.WorkspaceUID(workspaceUID),
			entweblink.DeactivatedAtIsNil(),
		).Count(ctx)
		return err
	})
	g.Go(func() (err error) {
		workspace, err = app.EntClient.Workspace.Query().Where(
			entworkspace.UID(workspaceUID),
			entworkspace.DeactivatedAtIsNil(),
		).First(ctx)
		return err
	})
	if err = g.Wait(); err != nil {
		return err
	}
	if app.IsNotBlank(title) {
		info.Host = title
	}
	if app.IsBlank(info.Host) {
		info.Host = "无效链接"
	}
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.WebLink.Create().
			SetTitle(info.Host).
			SetDescription(info.Title).
			SetImage(info.Image).
			SetLink(url).
			SetSequence(count).
			SetUserUID(userUID).
			SetFileType(app.URLFileType).
			SetWorkspaceUID(workspaceUID).
			SetFolderUID(workspace.PersonalFolderUID).
			Exec(ctx); err != nil {
			return err
		}
		if err = tx.Account.Update().AddURLCount(1).
			Where(entaccount.UID(userUID)).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

package article

import (
	"context"
	"golang.org/x/sync/errgroup"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/gen/entschema/goodarticle"
	entweblink "tabelf/backend/gen/entschema/weblink"
	"tabelf/backend/service/api/internal/logic/base"
	"tabelf/backend/service/api/internal/logic/station"
	"tabelf/backend/service/app"

	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodArticleCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodArticleCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodArticleCollectionLogic {
	return &GoodArticleCollectionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodArticleCollectionLogic) GoodArticleCollection(req *types.GoodArticleCollectionRequest) (resp *types.GoodArticleCollectionResponse, err error) {
	if err = base.ValidateUserAuthority(l.ctx, req.UserUID); err != nil {
		return nil, err
	}
	// 添加到该文件夹下到第一个工作空间，如果不存在给他创建一个
	workspace, err := station.GetFolderByFirstWorkspace(l.ctx, req.UserUID, req.FolderUID)
	if err != nil {
		return nil, err
	}
	var (
		g errgroup.Group
		//account *entschema.Account
		article *entschema.GoodArticle
		count   int
	)
	g.Go(func() (err error) {
		_, err = app.EntClient.Account.Query().Where(
			entaccount.UID(req.UserUID),
			entaccount.DeactivatedAtIsNil(),
		).First(l.ctx)
		if err != nil {
			if entschema.IsNotFound(err) {
				return app.ErrAccountInvalid(l.ctx)
			}
			return err
		}
		// 会员不做数量校验
		//if base.MembershipValidity(account) {
		//	return nil
		//}
		//if account.URLCount+1 > account.URLLimit {
		//	return app.ErrCustomerWebLinkLimit(l.ctx)
		//}
		return nil
	})
	g.Go(func() (err error) {
		article, err = app.EntClient.GoodArticle.Query().Where(
			goodarticle.UID(req.ArticleUID),
			goodarticle.DeactivatedAtIsNil(),
		).First(l.ctx)
		return err
	})
	g.Go(func() (err error) {
		count, err = app.EntClient.WebLink.Query().Where(
			entweblink.WorkspaceUID(workspace.UID),
			entweblink.DeactivatedAtIsNil(),
		).Count(l.ctx)
		return err
	})
	if err = g.Wait(); err != nil {
		return nil, err
	}
	if err = app.WithTx(l.ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if err = tx.WebLink.Create().
			SetTitle(article.Link).
			SetDescription(article.Description).
			SetImage(article.Icon).
			SetLink(article.Link).
			SetSequence(count).
			SetUserUID(req.UserUID).
			SetFileType(app.URLFileType).
			SetWorkspaceUID(workspace.UID).
			SetFolderUID(req.FolderUID).
			Exec(l.ctx); err != nil {
			return err
		}
		if err = tx.Account.Update().AddURLCount(1).
			Where(entaccount.UID(req.UserUID)).
			Exec(l.ctx); err != nil {
			return err
		}
		if _, err = base.UpdateGoodArticleMeta(
			l.ctx,
			req.ArticleUID,
			req.UserUID,
			app.ArticleStar,
		); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &types.GoodArticleCollectionResponse{
		Message: app.HttpOK,
	}, nil
}

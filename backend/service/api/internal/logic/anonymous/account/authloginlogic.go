package account

import (
	"context"
	entinvite "tabelf/backend/gen/entschema/invite"
	entpersonalfolder "tabelf/backend/gen/entschema/personalfolder"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"tabelf/backend/gen/entschema"
	entaccount "tabelf/backend/gen/entschema/account"
	"tabelf/backend/service/api/internal/svc"
	"tabelf/backend/service/api/internal/types"
	"tabelf/backend/service/api/pkg/utils"
	"tabelf/backend/service/app"
)

type AuthLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthLoginLogic {
	return &AuthLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthLoginLogic) AuthLogin(req *types.AuthLoginRequest) (resp *types.AuthLoginResponse, err error) {
	if app.IsBlank(req.AuthCode) {
		return nil, app.ErrAccountAuthCodeEmpty(l.ctx)
	}
	if len(req.AuthCode) != app.AuthCodeLen {
		return nil, app.ErrAccountAuthCodeInvalid(l.ctx)
	}
	account, err := app.EntClient.Account.Query().Where(
		entaccount.AuthCode(time.Now().Format("20060102") + req.AuthCode),
	).First(l.ctx)
	if err != nil {
		if entschema.IsNotFound(err) {
			return nil, app.ErrAccountAuthCodeNotExist(l.ctx)
		}
		return nil, err
	}
	if account.DeactivatedAt != nil {
		return nil, app.ErrAccountInvalid(l.ctx)
	}
	if time.Now().After(account.AuthExpired) {
		return nil, app.ErrAccountAuthCodeExpired(l.ctx)
	}
	// 如果为新用户
	if account.HasNew {
		if app.IsNotBlank(req.ReferralUID) { // 推荐奖励
			go func() {
				ctx := context.Background()
				if err1 := ReferralReward(
					ctx,
					req.ReferralUID,
					account.UID,
				); err1 != nil {
					app.Log.Error(ctx, err)
					return
				}
				app.Log.UnionLogger.Infof("pull new reward success, referral_uid = %+v, referee_uid = %+v", req.ReferralUID, account.UID)
			}()
		}
		// 默认绑定一个工作空间
		if err = BindDefaultWorkspace(l.ctx, account.UID); err != nil {
			return nil, err
		}
		// 变为老用户
		if err = app.EntClient.Account.Update().
			SetHasNew(false).
			AddURLCount(1).
			Where(entaccount.ID(account.ID)).
			Exec(l.ctx); err != nil {
			return nil, err
		}
	}
	JWTClaims := utils.JWTClaims{
		UID: account.UID,
	}
	JWTClaims.Audience = utils.AudienceUser
	token, err := utils.GetJwtToken(l.svcCtx.Config.Jwt.JwtKey, l.svcCtx.Config.Jwt.JwtIssuer,
		l.svcCtx.Config.Jwt.JwtExpire, JWTClaims)
	if err != nil {
		return nil, app.ErrAccountLoginFail(l.ctx)
	}
	return &types.AuthLoginResponse{
		UserUID:  account.UID,
		Username: account.Nickname,
		Token:    token,
	}, nil
}

func BindDefaultWorkspace(ctx context.Context, userUID string) (err error) {
	folderNumber := ""
	for {
		folderNumber = app.RandomString(app.FolderNumberCount)
		exist, err := app.EntClient.PersonalFolder.Query().Where(
			entpersonalfolder.FolderNumber(folderNumber),
			entpersonalfolder.UserUID(userUID),
			entpersonalfolder.DeactivatedAtIsNil(),
		).Exist(ctx)
		if err != nil {
			return err
		}
		if !exist {
			break
		}
	}
	var folder *entschema.PersonalFolder
	// 创建个人文件，并创建该文件对应的分享链接
	if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
		if folder, err = tx.PersonalFolder.Create().
			SetFolderName(app.ExampleFolderName).
			SetUserUID(userUID).
			SetFolderNumber(folderNumber).
			Save(ctx); err != nil {
			return err
		}
		workspace, err := tx.Workspace.Create().
			SetName(app.ExampleWorkspaceName).
			SetType(app.PersonalWorkspaceType).
			SetPersonalFolderUID(folder.UID).
			SetUserUID(userUID).
			SetIsOpen(true).
			Save(ctx)
		if err != nil {
			return err
		}
		if err = tx.WebLink.Create().
			SetTitle("www.zhihu.com").
			SetDescription("知乎 - 有问题，就会有答案").
			SetImage("https://static.zhihu.com/heifetz/favicon.ico").
			SetLink("https://www.zhihu.com").
			SetSequence(0).
			SetUserUID(userUID).
			SetFileType(app.URLFileType).
			SetWorkspaceUID(workspace.UID).
			SetFolderUID(workspace.PersonalFolderUID).
			Exec(ctx); err != nil {
			return err
		}
		return tx.ShareLink.Create().
			SetFolderUID(folder.UID).
			SetUserUID(userUID).
			SetAuthority(app.ShareReadAuthority).
			SetValidDay(app.ForEverDay).
			SetExpiredAt(app.ParseTime(app.ForEverValid)).
			SetRecentAt(time.Now()).
			SetFolderNumber(folder.FolderNumber).
			Exec(ctx)
	}); err != nil {
		return err
	}
	return nil
}

func ReferralReward(ctx context.Context, referralUID string, refereeUID string) (err error) {
	if refereeUID == referralUID {
		return nil
	}
	exist, err := app.EntClient.Invite.Query().Where(
		entinvite.ReferralUID(referralUID),
		entinvite.RefereeUID(refereeUID),
	).Exist(ctx)
	if err != nil {
		return err
	}
	account, err := app.EntClient.Account.Query().Where(
		entaccount.UID(referralUID),
		entaccount.DeactivatedAtIsNil(),
	).First(ctx)
	if err != nil {
		return err
	}
	if !exist {
		if err = app.WithTx(ctx, app.EntClient, func(tx *entschema.Tx) (err error) {
			if err = tx.Invite.Create().
				SetReferralUID(referralUID).
				SetRefereeUID(refereeUID).
				Exec(ctx); err != nil {
				return err
			}
			if err = tx.Account.Update().
				AddURLLimit(app.InviteRewardFileNum).
				Where(entaccount.ID(account.ID)).
				Exec(ctx); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}
	}
	return nil
}

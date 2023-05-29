package app

import (
	"context"
	"fmt"
	"tabelf/backend/common"
)

/* ======================== Account 错误定义 100000-100199 ========================. */
var (
	ErrAccountAuthCodeEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100000, "验证码为空")
	}
	ErrAccountAuthCodeInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100001, "验证码无效")
	}
	ErrAccountAuthCodeExpired = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100002, "验证码已过期")
	}
	ErrAccountAuthCodeNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100003, "验证码不存在")
	}

	ErrAccountInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100020, "用户信息无效")
	}
	ErrAccountLoginFail = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100021, "用户登录失败")
	}
	ErrAccountOutOfAuthority = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100022, "超出权限范围")
	}
	ErrAccountEmailInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100023, "无效的邮箱地址")
	}
	ErrAccountAdminNotAuthority = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100024, "用户无管理权限")
	}
	ErrAccountAdminNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100025, "管理员不存在")
	}
)

/* ======================== Folder 错误定义 100200-100299 ========================. */
var (
	ErrPersonalFolderNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000200, "文件夹不存在")
	}
	ErrPersonalFolderNameEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000201, "文件夹名称为空")
	}
	ErrPersonalFolderContentNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000202, "文件夹内容不存在")
	}
	ErrPersonalFolderNameLimit = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000203, "文件夹名称过长")
	}
)

/* ======================== Workspace 错误定义 100300-100399 ========================. */

/* ======================== WebLink 错误定义 100400-100499 ========================. */

/* ======================== Share 错误定义 100500-100599 ========================. */
var (
	ErrPersonalFolderShareLinkNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100500, "文件分享链接不存在")
	}
)

/* ======================== Share 错误定义 100600-100699 ========================. */
var (
	ErrPersonalFolderCollNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100600, "协作链接不存在")
	}
)

/* ======================== Feedback 错误定义 100700-100799 ========================. */
var (
	ErrPersonalFeedbackCategoryNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 100700, "反馈分类不存在")
	}
)

/* ======================== Customer 错误定义 100200-100499 ========================. */
var (
	ErrCustomerFolderShareAuthorityInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000202, "文件分享权限无效")
	}
	ErrCustomerFolderShareExpiredDayInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000203, "文件分享过期时间无效")
	}
	ErrCustomerFolderShareLinkNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000203, "文件分享链接不存在")
	}
	ErrCustomerFolderShareLinkExpired = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000203, "文件分享链接已过期")
	}

	ErrCustomerWorkspaceNameEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000300, "书签栏名称为空")
	}

	ErrCustomerWebLinkEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000401, "URL为空")
	}
	ErrCustomerWebLinkInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000402, "URL无效")
	}
	ErrCustomerWebLinkLimit = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000403, "链接数量超过限制")
	}
	ErrCustomerLocalFileLimit = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000404, "文件大小不能超过10M")
	}
	ErrCustomerLocalFileEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000405, "文件内容不能为空")
	}

	ErrCustomerShareOutOfEditAuthority = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000501, "用户没有更新权限")
	}

	ErrCustomerUIDEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000601, "用户UID不能为空")
	}
	ErrCustomerHeaderImageEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000602, "用户头像不能为空")
	}
	ErrCustomerHeaderImageTypeInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000603, "上传图片格式不正确")
	}
	ErrCustomerHeaderImageLimit = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000604, "上传图片大小不能超过2M")
	}
	ErrCustomerOrderInfoNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000605, "订单信息不存在")
	}
	ErrCustomerOrderTransactionNumberEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000606, "订单交易单号为空")
	}
	ErrCustomerOrderEventCanNot = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000607, "订单无法进行该操作")
	}
	ErrCustomerOrderUpgradeUIDEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000608, "订单会员权益消息为空")
	}
	ErrCustomerOrderTotalPriceEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000609, "订单总金额消息为空")
	}
	ErrCustomerOrderPaymentAmountEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000610, "订单实付金额消息为空")
	}
	ErrCustomerOrderPaymentTypeEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000611, "订单支付方式为空")
	}
	ErrCustomerOrderPaymentTypeInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000612, "订单支付方式无效")
	}
	ErrCustomerOrderPaymentAmountIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000613, "订单实付金额异常")
	}
	ErrCustomerOrderTotalPriceIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000614, "订单总金额异常")
	}
	ErrCustomerOrderPaymentRequest = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000615, "创建支付单失败")
	}
	ErrCustomerOrderPaymentRequestReason = func(ctx context.Context, msg string) *common.ZError {
		return NewError(ctx, 1000616, "创建支付单失败, 原因: "+msg)
	}
	ErrCustomerOrderNotExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000617, "订单消息不存在")
	}
	ErrCustomerOrderNotUpdate = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000618, "订单消息更新失败")
	}
	ErrCustomerUpdate = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000619, "用户消息更新失败")
	}

	ErrOrderNotifyIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000701, "订单回调异常")
	}
	ErrOrderNotifyCodeEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000702, "订单回调支付结果为空")
	}
	ErrOrderNotifyOrderNoEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000703, "订单回调系统订单号为空")
	}
	ErrOrderNotifyOutTradeNoEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000704, "订单回调商户订单号为空")
	}
	ErrOrderNotifyPayNoEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000705, "订单回调支付单号为空")
	}
	ErrOrderNotifyMoneyEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000706, "订单回调支付金额为空")
	}
	ErrOrderNotifyMchIDEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000707, "订单回调商户ID为空")
	}
	ErrOrderNotifyCodeIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000708, "订单回调支付结果无效")
	}
	ErrOrderNotifySignEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000709, "订单回调数字签名为空")
	}
	ErrOrderNotifySignIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000710, "订单回调数字签名无效")
	}
	ErrOrderNotifyOutTradeIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000711, "订单回调商户订单号无效")
	}
	ErrOrderNotifyAttachIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000712, "订单回调附加参数无效")
	}
	ErrOrderNotifyOrderStatusIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000713, "订单回调订单状态无效")
	}
	ErrOrderNotifyMoneyIllegal = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000714, "订单回调金额无效")
	}
)

var (
	ErrStationCategoryEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000801, "推荐分类不能为空")
	}
	ErrStationLinkEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000802, "推荐网站链接不能为空")
	}
	ErrStationLinkInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000803, "URL无效")
	}
	ErrStationSortedInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000804, "推荐排序查询无效")
	}
	ErrStationAudit = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000805, "好站推荐内容审核失败")
	}
)

var (
	ErrCommunityCategoryEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000901, "发布分类不能为空")
	}
	ErrCommunityTitleEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000902, "发布标题不能为空")
	}
	ErrCommunityDescriptionEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000903, "发布描述不能为空")
	}
	ErrCommunityFolderUIDEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000904, "发布文件不能为空")
	}
	ErrCommunityImageEmpty = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000905, "发布图片不能为空")
	}
	ErrCommunityImageParse = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000906, "发布图片解析失败")
	}
	ErrCommunitySortedInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000907, "排序查询无效")
	}
	ErrCommunityCategoryInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000908, "分类查询无效")
	}
	ErrCommunityAudit = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000909, "发布内容审核失败")
	}
	ErrCommunityRepeatPublic = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1000909, "资源已发布, 无需重复发布")
	}
)

var (
	ErrArticleSortedInvalid = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1001000, "排序查询无效")
	}
	ErrArticleExist = func(ctx context.Context) *common.ZError {
		return NewError(ctx, 1001001, "该文已存在")
	}
)

type Error struct {
	Code    string
	Message string
	TraceID string
	SpanID  string
}

func NewError(ctx context.Context, code interface{}, message string) *common.ZError {
	return &common.ZError{
		Code:    fmt.Sprintf("%v", code),
		Message: message,
		TraceID: common.TraceIDFromContext(ctx),
		SpanID:  common.SpanIDFromContext(ctx),
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

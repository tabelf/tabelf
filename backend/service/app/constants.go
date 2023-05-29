package app

import "time"

const (
	DataFormat = "2006年01月02日 15:04"

	GrpcAPMType   = "grpc"
	HTTPAPMType   = "http"
	HTTPProtoType = "http://"

	HttpOK          = "OK"
	HttpNO          = "NO"
	HttpSuccessCode = 200
)

const (
	DateWithTimeLayout = "2006-01-02 15:04:05"
	YMDLayout          = "2006-01-02"
)

const (
	AuthCodeLen       = 6
	AuthCodeExpiredAt = 5 * time.Minute // 5 分钟

	RecentWebLinkMaxLimitNum = 20

	UserOrderKey         = "UserOrder"
	TransactionNumberKey = "TransactionNumber"

	WechatPaymentType = "wechat"
	AlipayPaymentType = "alipay"

	OrderExpiredAt = 15 * time.Minute

	InviteRewardFileNum = 5 // 邀请奖励文件数 5

	Show   = true
	Hidden = false

	StationView   = 0
	StationPraise = 1
	StationStar   = 2
	StationNew    = 3

	CommunityView   = 0
	CommunityPraise = 1
	CommunityStar   = 2
	CommunityNew    = 3
	CommunityUsed   = 4

	ArticleView = 0
	ArticleUsed = 1
	ArticleStar = 2
	ArticleNew  = 3
)

const (
	RechargeOrderType = "recharge" // 充值
)

// Wechat.
const (
	QrScene        = "TabElf"
	AccessTokenURL = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	TicketURL      = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s"
	UserInfoURL    = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"

	MsgEventType = "event"
	MsgTextType  = "text"

	AuthCodeMsg     = "验证码"
	MenuAuthCodeMsg = "登录验证"
	ScanEvent       = "SCAN"
	SubscribeEvent  = "subscribe"
)

// Workspace.
const (
	PersonalWorkspaceType = 0
	TeamWorkspaceType     = 1

	DefaultWorkspaceName = "未命名"
	ExampleWorkspaceName = "示例1"
)

// Account.
const (
	NormalUser      = "0" // 普通用户
	MonthMemberUser = "1" // 月度会员
	YearMemberUser  = "2" // 年度会员

	UnknownUserName  = "未命名用户"
	UnknownSuffix    = 5
	UnknownUserImage = "https://tabelf-1316709639.cos.ap-shanghai.myqcloud.com/unknown_head_image.png"

	UnknownSex = 0

	NormalUserURLLimit = 30 // 正常用户 URL 限制数量为 30

	EmailSSLPorts     = "25"
	EmailValidate     = "[\\w]+(\\.[\\w]+)*@[\\w]+(\\.[\\w])+"
	EmailShareSubject = "邀请协作-分享链接"

	MaxBufferUploadMemory          = 1024 * 1024 * 3  // 3 M
	MaxUploadHeaderImageLimit      = 1024 * 1024 * 2  // 2 M
	MaxLocalFileBufferUploadMemory = 1024 * 1024 * 12 // 12 M
	MaxLocalFileUploadMemoryLimit  = 1024 * 1024 * 10 // 10 M
)

// Folder.
const (
	ShareReadAuthority   = "0"
	ShareEditAuthority   = "1"
	ShareRemoveAuthority = "3"

	FolderNumberCount = 8
	ExampleFolderName = "演示示例"

	ForEverDay   = -1
	ForEverValid = "2099-12-31 23:59:59" // 过期时间永久有效

	CopyMark           = " (复制)"
	CopyPersonalFolder = 0 // 个人文件拷贝
	CopyCollaboration  = 1 // 协作文件拷贝
)

// TODO: 需要指定自己的 腾讯COS 地址
const (
	TxCosURL          = "https://tabelf-1316709639.cos.ap-shanghai.myqcloud.com"           // 用户的图像
	TxCosLocalFileURL = "https://tabelf-local-1316709639.cos.ap-shanghai.myqcloud.com"     // 本地文件录入的文件
	TxCosStationURL   = "https://tabelf-station-1316709639.cos.ap-shanghai.myqcloud.com"   // 好站推荐的图片
	TxCosCommunityURL = "https://tabelf-community-1316709639.cos.ap-shanghai.myqcloud.com" // 社区交流的图片
	TxCosArticleURL   = "https://tabelf-article-1316709639.cos.ap-shanghai.myqcloud.com"   // 好文发现的图片

	WechatPaymentURL = "https://api.pay.yungouos.com/api/pay/wxpay/nativePay"
)

// Feedback.
const (
	OrderFeedbackCategory        = "支付问题"
	BugFeedbackCategory          = "Bug问题"
	OptimizationFeedbackCategory = "系统优化"
	OtherFeedbackCategory        = "其他问题"

	OrderFeedbackSystemUsername = "473225193@qq.com"
	EmailOrderFeedbackSubject   = "支付问题"
)

// Messgae.
type MessageType string

const (
	SystemMessage     MessageType = "system"     // 系统消息
	MembershipMessage MessageType = "membership" // 会员消息
	ShareMessage      MessageType = "share"      // 分享消息
	InviteMessage     MessageType = "invite"     // 邀请消息

	SystemAccount = "dd372d8c904e45268d198c1430b40410"

	MessageUnRead = 0 // 未读
	AllMessage    = 1 // 全部消息
)

// WebLink
const (
	URLFileType = "url" // 网络链接
)

// Community
const (
	FailPassStatus  = "-1" // 不通过
	WaitAuditStatus = "0"  // 待审核
	PassStatus      = "1"  // 审核通过

	SelfPublicCategory = "1" // 我的发布
	SelfStarCategory   = "2" // 我的收藏
	SelfRecentCategory = "3" // 最近使用
	SelfAuditCategory  = "4" // 管理审核
)

// GoodArticle
const HotArticleNum = 10

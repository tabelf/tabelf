// Code generated by goctl. DO NOT EDIT.
package types

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
}

type EmptyReq struct {
}

type ErrorReq struct {
	Code int `form:"code"`
}

type TimeResp struct {
	Now string `json:"now"`
}

type ErrorResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type HelloRequest struct {
	Name string `path:"name"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

type GetAuthLoginQrCodeRequest struct {
}

type GetAuthLoginQrCodeResponse struct {
	ExpiredAt string `json:"expired_at"` // 过期时间
	URL       string `json:"url"`        // 登录二维码
}

type AuthLoginRequest struct {
	AuthCode    string `json:"auth_code"`
	ReferralUID string `json:"referral_uid,optional"`
}

type AuthLoginResponse struct {
	UserUID  string `json:"user_uid"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetAuthLoginWechatAuthorityRequest struct {
	Signature string `form:"signature"`
	Nonce     string `form:"nonce"`
	Echostr   string `form:"echostr"`
	Timestamp string `form:"timestamp"`
}

type GetAuthLoginWechatAuthorityResponse struct {
	Echostr string `json:"echostr"`
}

type WechatMessageCallbackRequest struct {
	ToUserName   string `json:"ToUserName"`
	FromUserName string `json:"FromUserName"`
	CreateTime   string `json:"CreateTime"`
	MsgType      string `json:"MsgType"`
	Content      string `json:"Content"`
	Event        string `json:"Event"`
	EventKey     string `json:"EventKey"`
}

type WechatMessageCallbackResponse struct {
	Message string `json:"message"`
}

type GetGoodStationCategoryRequest struct {
	UserUID string `form:"user_uid,optional"`
}

type GetGoodStationCategoryResponse struct {
	Wait                  int                    `json:"wait"`
	HasAuthority          bool                   `json:"has_authority"`
	GoodStationCategories []*GoodStationCategory `json:"categories"`
}

type GoodStationCategory struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type GetGoodStationRecommendRequest struct {
	UserUID     string `form:"user_uid,optional"`
	CategoryUID string `form:"category_uid,optional"`
	Sorted      int    `form:"sorted"`
	Offset      int    `form:"offset"`
	Limit       int    `form:"limit"`
}

type GetGoodStationRecommendResponse struct {
	Total int                     `json:"total"`
	Data  []*GoodStationRecommend `json:"data"`
}

type GoodStationRecommend struct {
	UID          string   `json:"uid"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Image        string   `json:"image"`
	Tags         []string `json:"tags"`
	Icon         string   `json:"icon"`
	Source       string   `json:"source"`
	Link         string   `json:"link"`
	Praise       int      `json:"praise"`
	HasPraise    bool     `json:"has_praise"`
	Star         int      `json:"star"`
	HasStar      bool     `json:"has_star"`
	View         int      `json:"view"`
	CategoryName string   `json:"category_name"`
}

type GetPublicCommunityDetailRequest struct {
	CommunityUID string `path:"community_uid"`
	UserUID      string `form:"user_uid,optional"`
}

type GetPublicCommunityDetailResponse struct {
	UID         string   `json:"uid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	HtmlDesc    string   `json:"html_desc"`
	Image       string   `json:"image"`
	Tags        []string `json:"tags"`
	Praise      int      `json:"praise"`
	HasPraise   bool     `json:"has_praise"`
	Used        int      `json:"used"`
	View        int      `json:"view"`
	Star        int      `json:"star"`
	HasStar     bool     `json:"has_star"`
	UserUID     string   `json:"user_uid"`   // 用户uid
	UserImage   string   `json:"user_image"` // 用户图片
	UserName    string   `json:"user_name"`
	Fans        int      `json:"fans"`
	Open        int      `json:"open"`
	HasSelf     bool     `json:"has_self"`
	HasFollow   bool     `json:"has_follow"` // 是否关注
}

type GetPublicCommunityCategoryRequest struct {
	UserUID string `form:"user_uid,optional"`
}

type GetPublicCommunityCategoryResponse struct {
	Wait                    int                        `json:"wait"`
	HasAdmin                bool                       `json:"has_admin"`
	PublicCommunityCategory []*PublicCommunityCategory `json:"categories"`
}

type PublicCommunityCategory struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type GetPublicCommunityRequest struct {
	CategoryUID string `form:"category_uid,optional"`
	Sorted      int    `form:"sorted"`
	Offset      int    `form:"offset"`
	Limit       int    `form:"limit"`
}

type GetPublicCommunityResponse struct {
	Total int                `json:"total"`
	Data  []*PublicCommunity `json:"data"`
}

type PublicCommunity struct {
	UID         string `json:"uid"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Praise      int    `json:"praise"`
	Used        int    `json:"used"`
	View        int    `json:"view"`
	UserUID     string `json:"user_uid"`   // 用户uid
	UserImage   string `json:"user_image"` // 用户图片
	UserName    string `json:"user_name"`
	Status      string `json:"status"`
}

type ViewGoodArticleRequest struct {
	ArticleUID string `path:"article_uid"`
}

type ViewGoodArticleResponse struct {
	Message string `json:"message"`
}

type GetHotGoodArticleRequest struct {
}

type GetHotGoodArticleResponse struct {
	Data []*HotGoodArticle `json:"data"`
}

type HotGoodArticle struct {
	UID   string `json:"uid"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

type GetGoodArticleMenuDataRequest struct {
	UserUID string `form:"user_uid,optional"`
}

type GetGoodArticleMenuDataResponse struct {
	HasAnonymous bool `json:"has_anonymous"`
	HasAuthority bool `json:"has_authority"`
	Collection   int  `json:"collection"`
	Publish      int  `json:"publish"`
	Audit        int  `json:"audit"`
}

type GetGoodArticleCategoryRequest struct {
}

type GetGoodArticleCategoryResponse struct {
	GoodArticleCategories []*GoodArticleCategory `json:"categories"`
}

type GoodArticleCategory struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

type GetGoodArticleRecommendRequest struct {
	UserUID     string `form:"user_uid,optional"`
	CategoryUID string `form:"category_uid,optional"`
	Sorted      int    `form:"sorted"`
	Offset      int    `form:"offset"`
	Limit       int    `form:"limit"`
}

type GetGoodArticleRecommendResponse struct {
	Total int                     `json:"total"`
	Data  []*GoodArticleRecommend `json:"data"`
}

type GoodArticleRecommend struct {
	UID         string `json:"uid"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Source      string `json:"source"`
	Link        string `json:"link"`
	Star        int    `json:"star"`
	HasStar     bool   `json:"has_star"`
	Status      string `json:"status"`
}

type FocusPersonalAccountRequest struct {
	UserUID     string `path:"user_uid"`
	FolloweeUID string `json:"followee_uid"` // 被关注者
	Status      bool   `json:"status"`       // true 关注
}

type FocusPersonalAccountResponse struct {
	Message string `json:"message"`
}

type GetPersonalInviteRequest struct {
	UserUID string `path:"user_uid"`
}

type GetPersonalInviteResponse struct {
	InviteURL    string         `json:"invite_url"`   // 邀请链接
	Earned       int            `json:"earned"`       // 挣得文件数
	Invited      int            `json:"invited"`      // 邀请人数
	Achievements []*Achievement `json:"achievements"` // 成就, 最多20个，时间倒序
}

type Achievement struct {
	UserUID   string `json:"user_uid"` // 用户uid
	Image     string `json:"image"`    // 用户图片
	UserName  string `json:"user_name"`
	CreatedAt string `json:"created_at"`
}

type SearchWebLinksRequest struct {
	UserUID string `path:"user_uid"`
	Keyword string `form:"keyword,optional"`
	Type    int    `form:"type"` // 0 为全部, 1 为我的, 2 为协作
}

type SearchWebLinksResponse struct {
	SearchWebLinks []*SearchWebLink `json:"search_web_links"`
}

type SearchWebLink struct {
	FolderUID    string `json:"folder_uid"`
	FolderNumber string `json:"folder_number"`
	LinkUID      string `json:"link_uid"`
	Title        string `json:"title"`
	Image        string `json:"image"`
	Link         string `json:"link"`
	Description  string `json:"description"`
}

type UpdatePersonalAccountImageRequest struct {
	UserUID  string `path:"user_uid"`
	File     []byte `json:"file"`
	Filename string `json:"filename"`
}

type UpdatePersonalAccountImageResponse struct {
	ImageURL string `json:"image_url"`
}

type GetPersonalAccountInfoRequest struct {
	UserUID string `path:"user_uid"`
}

type GetPersonalAccountInfoResponse struct {
	UserUID       string `json:"user_uid"`
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	Industry      string `json:"industry"`
	Image         string `json:"image"`
	Description   string `json:"description"`
	HasMembership bool   `json:"has_membership"`
	HasEntire     bool   `json:"has_entire"` // 是否完成了基本信息
}

type UpdatePersonalAccountInfoRequest struct {
	UserUID     string `path:"user_uid"`
	UserName    string `json:"user_name,optional"`
	Email       string `json:"email,optional"`
	Industry    string `json:"industry,optional"`
	Description string `json:"description,optional"`
}

type UpdatePersonalAccountInfoResponse struct {
	Message string `json:"message"`
}

type GetPersonalWebLinkInfoRequest struct {
	UserUID string `path:"user_uid"`
}

type GetPersonalWebLinkInfoResponse struct {
	UsedQuantity  int  `json:"used_quantity"`
	TotalQuantity int  `json:"total_quantity"`
	Percent       int  `json:"percent"`
	HasLimit      bool `json:"has_limit"`
}

type GetPersonalTeamsRequest struct {
	UserUID string `path:"user_uid"`
}

type GetPersonalTeamsResponse struct {
	PersonalTeams []*PersonalTeam `json:"personal_teams"`
}

type PersonalTeam struct {
	TeamUID   string `json:"team_uid"`
	TeamName  string `json:"team_name"`
	ExpiredAt string `json:"expired_at"`
}

type CopyPersonalFolderRequest struct {
	UserUID   string `path:"user_uid"`
	FolderUID string `path:"folder_uid"`
	Type      int    `json:"type"`
}

type CopyPersonalFolderResponse struct {
	Message string `json:"message"`
}

type GetPersonalFoldersRequest struct {
	UserUID string `path:"user_uid"`
}

type GetPersonalFoldersResponse struct {
	Folders []*PersonalFolder `json:"folders"`
}

type PersonalFolder struct {
	FolderUID    string `json:"folder_uid"`
	FolderNumber string `json:"folder_number"`
	FolderName   string `json:"folder_name"`
}

type AddPersonalFolderRequest struct {
	UserUID    string `path:"user_uid"`
	FolderName string `json:"folder_name"`
}

type AddPersonalFolderResponse struct {
	FolderUID    string `json:"folder_uid"`
	FolderNumber string `json:"folder_number"`
}

type UpdatePersonalFolderRequest struct {
	UserUID    string `path:"user_uid"`
	FolderUID  string `path:"folder_uid"`
	FolderName string `json:"folder_name"`
}

type UpdatePersonalFolderResponse struct {
	Message string `json:"message"`
}

type DeletePersonalFolderRequest struct {
	UserUID   string `path:"user_uid"`
	FolderUID string `path:"folder_uid"`
}

type DeletePersonalFolderResponse struct {
	Message string `json:"message"`
}

type GetWorkspaceContentRequest struct {
	UserUID   string `path:"user_uid"`
	FolderUID string `path:"folder_uid"`
}

type GetWorkspaceContentResponse struct {
	Authority           string               `json:"authority"`
	ShareUID            string               `json:"share_uid"`
	FolderUID           string               `json:"folder_uid"`
	FolderNumber        string               `json:"folder_number"`
	FolderName          string               `json:"folder_name"`
	ActiveWorkspaceUIDs []string             `json:"active_workspace_uids"`
	PersonalWorkspaces  []*PersonalWorkspace `json:"personal_workspaces"`
}

type GetFolderNumberWorkspaceContentRequest struct {
	UserUID      string `path:"user_uid"`
	FolderNumber string `form:"folder_number"`
}

type GetFolderNumberWorkspaceContentResponse struct {
	FolderUID           string               `json:"folder_uid"`
	FolderNumber        string               `json:"folder_number"`
	FolderName          string               `json:"folder_name"`
	ActiveWorkspaceUIDs []string             `json:"active_workspace_uids"`
	PersonalWorkspaces  []*PersonalWorkspace `json:"personal_workspaces"`
}

type AddWorkspaceRequest struct {
	UserUID       string `path:"user_uid"`
	FolderUID     string `path:"folder_uid"`
	WorkspaceName string `json:"workspace_name"`
}

type AddWorkspaceResponse struct {
	Message string `json:"message"`
}

type UpdateWorkspaceRequest struct {
	UserUID       string `path:"user_uid"`
	FolderUID     string `path:"folder_uid"`
	WorkspaceUID  string `path:"workspace_uid"`
	WorkspaceName string `json:"workspace_name"`
}

type UpdateWorkspaceResponse struct {
	Message string `json:"message"`
}

type DeleteWorkspaceRequest struct {
	UserUID      string `path:"user_uid"`
	FolderUID    string `path:"folder_uid"`
	WorkspaceUID string `path:"workspace_uid"`
}

type DeleteWorkspaceResponse struct {
	Message string `json:"message"`
}

type UpdateWorkspaceSwitchRequest struct {
	UserUID             string   `path:"user_uid"`
	FolderUID           string   `path:"folder_uid"`
	ActiveWorkspaceUIDs []string `json:"active_workspace_uids"`
}

type UpdateWorkspaceSwitchResponse struct {
	Message string `json:"message"`
}

type GetFolderNumberShareWorkspaceContentRequest struct {
	UserUID      string `path:"user_uid"`
	FolderNumber string `form:"folder_number"`
}

type GetFolderNumberShareWorkspaceContentResponse struct {
	Authority           string               `json:"authority"`
	ShareUID            string               `json:"share_uid"`
	FolderUID           string               `json:"folder_uid"`
	FolderNumber        string               `json:"folder_number"`
	FolderName          string               `json:"folder_name"`
	ActiveWorkspaceUIDs []string             `json:"active_workspace_uids"`
	PersonalWorkspaces  []*PersonalWorkspace `json:"personal_workspaces"`
}

type PersonalWorkspace struct {
	IsOpen        bool               `json:"is_open"`
	WorkspaceUID  string             `json:"workspace_uid"`
	WorkspaceName string             `json:"workspace_name"`
	WebLinks      []*PersonalWebLink `json:"web_links"`
}

type RestoreDeletePersonalWebLinkRequest struct {
	UserUID string `path:"user_uid"`
	LinkUID string `path:"link_uid"`
}

type RestoreDeletePersonalWebLinkResponse struct {
	Message string `json:"message"`
}

type DeleteForeverPersonalWebLinkRequest struct {
	UserUID string `path:"user_uid"`
	LinkUID string `path:"link_uid"`
}

type DeleteForeverPersonalWebLinkResponse struct {
	Message string `json:"message"`
}

type GetRecyclingPersonalWebLinkRequest struct {
	UserUID string `path:"user_uid"`
}

type GetRecyclingPersonalWebLinkResponse struct {
	WebLinks []*PersonalWebLink `json:"web_links"`
}

type GetRecentPersonalWebLinkRequest struct {
	UserUID string `path:"user_uid"`
}

type GetRecentPersonalWebLinkResponse struct {
	WebLinks []*RecentWebLink `json:"web_links"`
}

type RecentWebLink struct {
	LinkUID     string `json:"link_uid"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	FileType    string `json:"file_type"`
	Description string `json:"description"`
	UpdatedAt   string `json:"updated_at"`
}

type DeletePersonalWebLinkRequest struct {
	UserUID      string `path:"user_uid"`
	FolderUID    string `path:"folder_uid"`
	WorkspaceUID string `path:"workspace_uid"`
	LinkUID      string `path:"link_uid"`
}

type DeletePersonalWebLinkResponse struct {
	Message string `json:"message"`
}

type UpdatePersonalWebLinkRequest struct {
	UserUID      string `path:"user_uid"`
	FolderUID    string `path:"folder_uid"`
	WorkspaceUID string `path:"workspace_uid"`
	LinkUID      string `path:"link_uid"`
	Title        string `json:"title"`
	Description  string `json:"description"`
}

type UpdatePersonalWebLinkResponse struct {
	Message string `json:"message"`
}

type AddPersonalWebLinkRequest struct {
	UserUID      string `path:"user_uid"`
	FolderUID    string `path:"folder_uid"`
	WorkspaceUID string `path:"workspace_uid"`
	Title        string `json:"title,optional"`
	URL          string `json:"url"`
}

type AddPersonalLocalFileLinkRequest struct {
	UserUID      string `path:"user_uid"`
	FolderUID    string `path:"folder_uid"`
	WorkspaceUID string `path:"workspace_uid"`
	File         []byte `json:"file,optional"`
	Filename     string `json:"filename,optional"`
	FileType     string `json:"file_type,optional"`
}

type AddPersonalLocalFileLinkResponse struct {
	Message string `json:"message"`
}

type AddPersonalWebLinkResponse struct {
	Message string `json:"message"`
}

type UpdateWorkspaceWebLinksRequest struct {
	UserUID         string             `path:"user_uid"`
	FolderUID       string             `path:"folder_uid"`
	WorkspaceUID    string             `path:"workspace_uid"`
	OldWorkspaceUID string             `json:"old_workspace_uid"`
	WebLinks        []*PersonalWebLink `json:"web_links"`
}

type UpdateWorkspaceWebLinksResponse struct {
	Message string `json:"message"`
}

type PersonalWebLink struct {
	LinkUID     string `json:"link_uid"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Link        string `json:"link"`
	FileType    string `json:"file_type"`
	Description string `json:"description"`
	Sequence    int    `json:"sequence"`
}

type AddSharePersonalWebLinkRequest struct {
	UserUID      string `path:"user_uid"`
	ShareUID     string `path:"share_uid"`
	WorkspaceUID string `path:"workspace_uid"`
	URL          string `json:"url"`
	Title        string `json:"title,optional"`
}

type AddSharePersonalWebLinkResponse struct {
	Message string `json:"message"`
}

type UpdateShareLinkRequest struct {
	UserUID    string `path:"user_uid"`
	FolderUID  string `path:"folder_uid"`
	ShareUID   string `path:"share_uid"`
	Authority  string `json:"authority"`
	ExpiredDay int    `json:"expired_day"`
}

type UpdateShareLinkResponse struct {
	Message string `json:"message"`
}

type JoinShareLinkRequest struct {
	UserUID  string `path:"user_uid"`
	ShareUID string `path:"share_uid"`
}

type JoinShareLinkResponse struct {
	HasShare     bool   `json:"has_share"` // 是否用于分享，false表示自己访问自己的链接
	FolderUID    string `json:"folder_uid"`
	FolderNumber string `json:"folder_number"`
}

type ShareFolderLinkRequest struct {
	UserUID    string `path:"user_uid"`
	FolderUID  string `path:"folder_uid"`
	Authority  string `json:"authority"`
	ExpiredDay int    `json:"expired_day"`
}

type ShareFolderLinkResponse struct {
	ShareUID   string `json:"share_uid"`
	ShareLink  string `json:"share_link"`
	Authority  string `json:"authority"`
	ExpiredDay int    `json:"expired_day"`
}

type ExitSharePersonalFolderRequest struct {
	UserUID  string `path:"user_uid"`
	ShareUID string `path:"share_uid"`
}

type ExitSharePersonalFolderResponse struct {
	Message string `json:"message"`
}

type ShareFriendFolderRequest struct {
	UserUID   string `path:"user_uid"`
	ShareUID  string `path:"share_uid"`
	Email     string `json:"email"`
	Authority string `json:"authority"`
}

type ShareFriendFolderResponse struct {
	Message string `json:"message"`
}

type UpdateSharePersonnelRequest struct {
	UserUID   string `path:"user_uid"`
	FolderUID string `path:"folder_uid"`
	ShareUID  string `path:"share_uid"`
	CollUID   string `path:"coll_uid"`
	Type      string `json:"type"` // 操作类型：0 为可查看，1 为可编辑，3 为删除
}

type UpdateSharePersonnelResponse struct {
	Message string `json:"message"`
}

type SharePersonalFolderRequest struct {
	UserUID   string `path:"user_uid"`
	FolderUID string `path:"folder_uid"`
	Offset    int    `form:"offset"`
	Limit     int    `form:"limit"`
}

type SharePersonalFolderResponse struct {
	FolderUID      string            `json:"folder_uid"`
	ShareUID       string            `json:"share_uid"`
	ShareLink      string            `json:"share_link"`
	Authority      string            `json:"authority"`
	ExpiredDay     int               `json:"expired_day"`
	HasOwner       bool              `json:"has_owner"` // 是否是创建者
	ShareNum       int               `json:"share_num"`
	SharePersonnel []*SharePersonnel `json:"share_personnel"`
}

type GetShareAllPersonnelRequest struct {
	UserUID  string `path:"user_uid"`
	ShareUID string `path:"share_uid"`
}

type GetShareAllPersonnelResponse struct {
	HasOwner       bool              `json:"has_owner"` // 是否是创建者
	SharePersonnel []*SharePersonnel `json:"share_personnel"`
}

type SharePersonnel struct {
	CollUID       string `json:"coll_uid"`
	UserUID       string `json:"user_uid"`
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	Authority     string `json:"authority"`
	HasMembership bool   `json:"has_membership"`
	Image         string `json:"image"`
	HasSelf       bool   `json:"has_self"`
	Sequence      int    `json:"sequence"`
}

type UpdateManageSharePersonnelRequest struct {
	UserUID      string `path:"user_uid"`
	ShareUID     string `path:"share_uid"`
	OtherUserUID string `json:"other_user_uid"`
	Authority    string `json:"authority"`
}

type UpdateManageSharePersonnelResponse struct {
	Message string `json:"message"`
}

type DeleteSharePersonnelRequest struct {
	UserUID      string `path:"user_uid"`
	ShareUID     string `path:"share_uid"`
	OtherUserUID string `json:"other_user_uid"`
}

type DeleteSharePersonnelResponse struct {
	Message string `json:"message"`
}

type GetSharePersonalFolderRequest struct {
	UserUID string `path:"user_uid"`
}

type GetSharePersonalFolderResponse struct {
	Folders []*ShareFolder `json:"folders"`
}

type ShareFolder struct {
	ShareUID     string `json:"share_uid"`
	HasOwner     bool   `json:"has_owner"`
	Authority    string `json:"authority"`
	FolderUID    string `json:"folder_uid"`
	FolderNumber string `json:"folder_number"`
	FolderName   string `json:"folder_name"`
}

type UpdateSharePersonalFolderRequest struct {
	UserUID    string `path:"user_uid"`
	ShareUID   string `path:"share_uid"`
	FolderName string `json:"folder_name"`
}

type UpdateSharePersonalFolderResponse struct {
	Message string `json:"message"`
}

type DeleteSharePersonalFolderRequest struct {
	UserUID  string `path:"user_uid"`
	ShareUID string `path:"share_uid"`
}

type DeleteSharePersonalFolderResponse struct {
	Message string `json:"message"`
}

type CopySharePersonalFolderRequest struct {
	UserUID  string `path:"user_uid"`
	ShareUID string `path:"share_uid"`
	Type     int    `json:"type"`
}

type CopySharePersonalFolderResponse struct {
	Message string `json:"message"`
}

type GetUpgradeRechargesRequest struct {
}

type GetUpgradeRechargesResponse struct {
	DefaultAmount    string              `json:"amount"`
	UpgradeRecharges []*UpgradeRecharges `json:"upgrade_recharges"`
}

type UpgradeRecharges struct {
	UID          string   `json:"uid"`
	Title        string   `json:"title"`
	OriginAmount string   `json:"origin_amount"`
	Amount       string   `json:"amount"`
	Descriptions []string `json:"descriptions"`
	ThemeColor   string   `json:"theme_color"`
}

type CreateFeedbackRequest struct {
	UserUID     string `path:"user_uid"`
	Category    string `json:"category"`
	Description string `json:"description,optional"`
	OrderNumber string `json:"order_number,optional"`
}

type CreateFeedbackResponse struct {
	Message string `json:"message"`
}

type DelAllMessageRequest struct {
	UserUID string `path:"user_uid"`
}

type DelAllMessageResponse struct {
	Message string `json:"message"`
}

type DelMessageRequest struct {
	UserUID    string `path:"user_uid"`
	MessageUID string `path:"message_uid"`
}

type DelMessageResponse struct {
	Message string `json:"message"`
}

type ReadAllMessageRequest struct {
	UserUID string `path:"user_uid"`
}

type ReadAllMessageResponse struct {
	Message string `json:"message"`
}

type GetMessagesRequest struct {
	UserUID string `path:"user_uid"`
	MsgType int    `form:"msg_type"` // 0 未读消息，1 全部消息
}

type GetMessagesResponse struct {
	Unread       int            `json:"unread"` // 未读数量
	UserMessages []*UserMessage `json:"user_messages"`
}

type UserMessage struct {
	UID           string `json:"uid"`
	PromoterName  string `json:"promoter_name"`  // 发起人名称
	PromoterImage string `json:"promoter_image"` // 发起人头像
	Description   string `json:"description"`
	HasRead       bool   `json:"has_read"`
	CreatedAt     string `json:"created_at"`
}

type ReadMessageRequest struct {
	UserUID    string `path:"user_uid"`
	MessageUID string `path:"message_uid"`
}

type ReadMessageResponse struct {
	Message string `json:"message"`
}

type UpdateGoodStationRequest struct {
	UserUID     string   `path:"user_uid"`
	StationUID  string   `path:"station_uid"`
	CategoryUID string   `json:"category_uid"`
	SiteName    string   `json:"site_name"` // 网站名称
	Link        string   `json:"link"`
	Title       string   `json:"title,optional"`
	Description string   `json:"description,optional"`
	Image       string   `json:"image,optional"`
	Tags        []string `json:"tags,optional"`
}

type UpdateGoodStationResponse struct {
	Message string `json:"message"`
}

type GetGoodStationRequest struct {
	UserUID    string `path:"user_uid"`
	StationUID string `path:"station_uid"`
}

type GetGoodStationResponse struct {
	StationUID  string   `json:"station_uid"`
	CategoryUID string   `json:"category_uid"`
	SiteName    string   `json:"site_name"` // 网站名称
	Link        string   `json:"link"`
	Title       string   `json:"title,optional"`
	Description string   `json:"description,optional"`
	Image       string   `json:"image,optional"`
	Tags        []string `json:"tags,optional"`
}

type GoodStationStarRequest struct {
	UserUID    string `path:"user_uid"`
	StationUID string `path:"station_uid"`
	FolderUID  string `json:"folder_uid"`
}

type GoodStationStarResponse struct {
	UID       string `json:"uid"`
	Praise    int    `json:"praise"`
	HasPraise bool   `json:"has_praise"`
	Star      int    `json:"star"`
	HasStar   bool   `json:"has_star"`
	View      int    `json:"view"`
}

type UpdateGoodStationMetaRequest struct {
	UserUID    string `path:"user_uid"`
	StationUID string `path:"station_uid"`
	MetaType   int    `json:"meta_type"`
}

type UpdateGoodStationMetaResponse struct {
	UID       string `json:"uid"`
	Praise    int    `json:"praise"`
	HasPraise bool   `json:"has_praise"`
	Star      int    `json:"star"`
	HasStar   bool   `json:"has_star"`
	View      int    `json:"view"`
}

type UploadStationImageRequest struct {
	UserUID  string `path:"user_uid"`
	File     []byte `json:"file"`
	Filename string `json:"filename"`
}

type UploadStationImageResponse struct {
	ImageURL string `json:"image_url"`
}

type AddGoodStationRecommendRequest struct {
	UserUID     string   `path:"user_uid"`
	CategoryUID string   `json:"category_uid"`
	SiteName    string   `json:"site_name"` // 网站名称
	Link        string   `json:"link"`
	Title       string   `json:"title,optional"`
	Description string   `json:"description,optional"`
	Image       string   `json:"image,optional"`
	Tags        []string `json:"tags,optional"`
}

type AddGoodStationRecommendResponse struct {
	Message string `json:"message"`
}

type UpdateAuditGoodStationRequest struct {
	UserUID    string `path:"user_uid"`
	StationUID string `path:"station_uid"`
	Status     bool   `json:"status"`
}

type UpdateAuditGoodStationResponse struct {
	Message string `json:"message"`
}

type GetAuditGoodStationRequest struct {
	UserUID string `path:"user_uid"`
}

type GetAuditGoodStationResponse struct {
	Data []*GoodStationRecommend `json:"data"`
}

type GetSelfPublicCommunityRequest struct {
	UserUID  string `path:"user_uid"`
	Category string `form:"category"`
}

type GetSelfPublicCommunityResponse struct {
	HasAdmin bool               `json:"has_admin"`
	Data     []*PublicCommunity `json:"data"`
}

type ClosePublicCommunityRequest struct {
	UserUID      string `path:"user_uid"`
	CommunityUID string `path:"community_uid"`
}

type ClosePublicCommunityResponse struct {
	Message string `json:"message"`
}

type UpdatePublicCommunityImageRequest struct {
	UserUID   string `path:"user_uid"`
	FolderUID string `json:"folder_uid"`
	Image     string `json:"image"`
}

type UpdatePublicCommunityImageResponse struct {
	Message string `json:"message"`
}

type UpdatePublicCommunityMetaRequest struct {
	UserUID      string `path:"user_uid"`
	CommunityUID string `path:"community_uid"`
	MetaType     int    `json:"meta_type"`
}

type UpdatePublicCommunityMetaResponse struct {
	UID          string `json:"uid"`
	FolderNumber string `json:"folder_number"`
	Praise       int    `json:"praise"`
	HasPraise    bool   `json:"has_praise"`
	Star         int    `json:"star"`
	HasStar      bool   `json:"has_star"`
	View         int    `json:"view"`
	Used         int    `json:"used"`
	HasUsed      bool   `json:"has_used"`
}

type CreatePublicCommunityRequest struct {
	UserUID     string   `path:"user_uid"`
	CategoryUID string   `json:"category_uid"`
	FolderUID   string   `json:"folder_uid"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,optional"`
}

type CreatePublicCommunityResponse struct {
	Message string `json:"message"`
}

type UpdateAuditPublicCommunityRequest struct {
	UserUID      string `path:"user_uid"`
	CommunityUID string `path:"community_uid"`
	Status       string `json:"status"` // 1 通过, -1 拒绝
}

type UpdateAuditPublicCommunityResponse struct {
	Status string `json:"status"`
}

type GetAuditPublicCommunityRequest struct {
	UserUID string `path:"user_uid"`
}

type GetAuditPublicCommunityResponse struct {
	CommunityUID         string            `json:"community_uid"`
	CommunityTitle       string            `json:"community_title"`
	CommunityDescription string            `json:"community_description"`
	Status               string            `json:"status"`
	FolderUID            string            `json:"folder_uid"`
	FolderNumber         string            `json:"folder_number"`
	FolderName           string            `json:"folder_name"`
	AuditCount           int               `json:"audit_count"`
	PersonalWorkspaces   []*AuditWorkspace `json:"personal_workspaces"`
}

type AuditWorkspace struct {
	IsOpen        bool               `json:"is_open"`
	WorkspaceUID  string             `json:"workspace_uid"`
	WorkspaceName string             `json:"workspace_name"`
	WebLinks      []*PersonalWebLink `json:"web_links"`
}

type GetGoodArticleSubmitRequest struct {
	UserUID string `path:"user_uid"`
}

type GetGoodArticleSubmitResponse struct {
	Data []*GoodArticleRecommend `json:"data"`
}

type UpdateGoodArticleAuditRequest struct {
	UserUID    string `path:"user_uid"`
	ArticleUID string `path:"article_uid"`
	Status     bool   `json:"status"`
}

type UpdateGoodArticleAuditResponse struct {
	Message string `json:"message"`
}

type GetGoodArticleAuditRequest struct {
	UserUID string `path:"user_uid"`
}

type GetGoodArticleAuditResponse struct {
	Data []*GoodArticleRecommend `json:"data"`
}

type GetGoodArticleCollectionRequest struct {
	UserUID string `path:"user_uid"`
}

type GetGoodArticleCollectionResponse struct {
	Data []*GoodArticleRecommend `json:"data"`
}

type GetGoodArticleRequest struct {
	UserUID    string `path:"user_uid"`
	ArticleUID string `path:"article_uid"`
}

type GetGoodArticleResponse struct {
	ArticleUID  string `json:"article_uid"`
	CategoryUID string `json:"category_uid"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Source      string `json:"source"`
	Image       string `json:"image"`
}

type UpdateGoodArticleRequest struct {
	UserUID     string `path:"user_uid"`
	ArticleUID  string `json:"article_uid"`
	CategoryUID string `json:"category_uid"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Source      string `json:"source"`
	Image       string `json:"image,optional"`
}

type UpdateGoodArticleResponse struct {
	UID         string `json:"uid"`
	CategoryUID string `json:"category_uid"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Source      string `json:"source"`
	Image       string `json:"image"`
	Status      string `json:"status"`
}

type AddGoodArticleRecommendRequest struct {
	UserUID     string `path:"user_uid"`
	CategoryUID string `json:"category_uid"`
	Title       string `json:"title,optional"`
	Link        string `json:"link"`
	Source      string `json:"source"`
	Image       string `json:"image,optional"`
}

type AddGoodArticleRecommendResponse struct {
	Message string `json:"message"`
}

type UploadGoodArticleImageRequest struct {
	UserUID  string `path:"user_uid"`
	File     []byte `json:"file"`
	Filename string `json:"filename"`
}

type UploadGoodArticleImageResponse struct {
	ImageURL string `json:"image_url"`
}

type UpdateGoodArticleMetaRequest struct {
	UserUID    string `path:"user_uid"`
	ArticleUID string `path:"article_uid"`
	MetaType   int    `json:"meta_type"`
}

type UpdateGoodArticleMetaResponse struct {
	UID     string `json:"uid"`
	HasStar bool   `json:"has_star"`
}

type GoodArticleCollectionRequest struct {
	UserUID    string `path:"user_uid"`
	ArticleUID string `path:"article_uid"`
	FolderUID  string `json:"folder_uid"`
}

type GoodArticleCollectionResponse struct {
	Message string `json:"message"`
}

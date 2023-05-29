package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"

	"tabelf/backend/common"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	common.BaseSchema
}

type Address struct {
	Country  string `json:"country"`  // 国家
	Province string `json:"province"` // 省份
	City     string `json:"city"`     // 城市
}

// Fields fields.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("phone").Optional().Comment("手机号"),
		field.String("email").Optional().Comment("电子邮件"),
		field.String("nickname").Optional().Comment("昵称"),
		field.Int("sex").Default(0).Comment("性别, 1 为男性, 2 为女性, 0 为未知"),
		field.String("password").Optional().Comment("密码"),
		field.String("wx_openid").Optional().Comment("微信openid").Unique(),
		field.String("image").Optional().Comment("头像"),
		field.String("industry").Optional().Comment("行业"),
		field.String("description").Optional().Comment("个人简介"),
		field.JSON("address", Address{}).Optional().Comment("地址信息"),
		field.String("auth_code").Comment("验证码").Unique(),
		field.Time("auth_expired").SchemaType(DatetimeSchema).Comment("验证码过期时间"),
		field.Int("url_count").Default(0).Comment("url数量"),
		field.Int("url_limit").Default(30).Comment("url数量限制"),
		field.String("member_type").Default("0").Comment("会员类型, 0 普通用户, 1 月度会员, 2 年度会员"),
		field.Time("member_expired").Optional().SchemaType(DatetimeSchema).Comment("会员过期时间"),
		field.Bool("has_entire").Default(false).Comment("信息是否完善"),
		field.Bool("has_new").Default(true).Comment("是否为新用户"),
		field.Bool("has_admin").Default(false).Comment("是否为管理员"),
		field.Int("fans").Default(0).Comment("粉丝数量"),
		field.Int("focus").Default(0).Comment("关注数量"),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Account) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account"},
	}
}

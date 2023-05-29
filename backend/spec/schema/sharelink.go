package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ShareLink holds the schema definition for the ShareLink entity.
type ShareLink struct {
	common.BaseSchema
}

// Fields fields.
func (ShareLink) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("folder_uid").Comment("文件uid").Unique(),
		field.String("user_uid").Comment("用户uid"),
		field.String("authority").Comment("权限: 0 只读权限, 1 编辑权限"),
		field.Int("valid_day").Comment("有效期, -1 为永久有效"),
		field.Time("expired_at").SchemaType(DatetimeSchema).Comment("链接过期时间"),
		field.Time("recent_at").SchemaType(DateSchema).Comment("最近更新时间"),
		field.String("folder_number").MaxLen(LenFolderNumber).MinLen(LenFolderNumber).Comment("文件编号"),
	}
}

// Edges of the Account.
func (ShareLink) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (ShareLink) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
		index.Fields("folder_number"),
		index.Fields("folder_number", "user_uid"),
	}
}

func (ShareLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "share_link"},
	}
}

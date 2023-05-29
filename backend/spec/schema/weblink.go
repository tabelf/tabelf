package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// WebLink holds the schema definition for the WebLink entity.
type WebLink struct {
	common.BaseSchema
}

// Fields fields.
func (WebLink) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("title").Comment("网页标题"),
		field.String("image").Comment("网页图片"),
		field.Text("link").Comment("网页链接"),
		field.String("description").Default("").Comment("链接描述"),
		field.String("file_type").Default("").Comment("文件类型"),
		field.Int("sequence").Default(0).Comment("序号"),
		field.Bool("forever_delete").Default(false).Comment("永久删除, 0 否 1 是"),
		field.String("user_uid").Comment("用户uid"),
		field.String("workspace_uid").Comment("工作空间uid"),
		field.String("folder_uid").Comment("文件uid"),
	}
}

// Edges of the Account.
func (WebLink) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (WebLink) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
		index.Fields("workspace_uid"),
	}
}

func (WebLink) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "web_link"},
	}
}

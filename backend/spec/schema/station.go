package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Station holds the schema definition for the Team entity.
type Station struct {
	common.BaseSchema
}

// Fields fields.
func (Station) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("title").Comment("标题"),
		field.String("description").SchemaType(VarcharDesc).Optional().Comment("描述"),
		field.String("image").Comment("图片"),
		field.JSON("tags", []string{}).Optional().Comment("标签列表"),
		field.String("icon").Comment("图标"),
		field.String("source").Comment("来源"),
		field.String("link").Comment("链接"),
		field.Int("praise").Default(0).Comment("点赞量"),
		field.Int("star").Default(0).Comment("收藏量"),
		field.Int("view").Default(0).Comment("查看量"),
		field.String("user_uid").Comment("用户uid"),
		field.Bool("status").Default(false).Comment("状态"),
		field.String("category_uid").Comment("分类uid"),
	}
}

// Edges of the Account.
func (Station) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Station) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Station) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "station"},
	}
}

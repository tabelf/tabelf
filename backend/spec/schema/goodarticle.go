package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GoodArticle holds the schema definition for the Team entity.
type GoodArticle struct {
	common.BaseSchema
}

// Fields fields.
func (GoodArticle) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("title").Comment("标题"),
		field.String("description").SchemaType(VarcharDesc).Optional().Comment("描述"),
		field.String("image").Optional().Comment("图片"),
		field.String("source").Comment("来源"),
		field.String("icon").Optional().Comment("来源icon"),
		field.String("link").Comment("链接"),
		field.Int("star").Default(0).Comment("收藏量"),
		field.Int("view").Default(0).Comment("查看量"),
		field.Int("used").Default(0).Comment("使用量"),
		field.String("user_uid").Comment("用户uid"),
		field.String("status").Default("0").Comment("状态, -1 审核失败, 0 待审核, 1 审核通过"),
		field.String("category_uid").Comment("分类uid"),
	}
}

// Edges of the GoodArticle.
func (GoodArticle) Edges() []ent.Edge {
	return nil
}

// Indexes of the GoodArticle.
func (GoodArticle) Indexes() []ent.Index {
	return []ent.Index{}
}

func (GoodArticle) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "good_article"},
	}
}

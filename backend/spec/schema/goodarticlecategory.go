package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GoodArticleCategory holds the schema definition for the Team entity.
type GoodArticleCategory struct {
	common.BaseSchema
}

// Fields fields.
func (GoodArticleCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Comment("分类名称"),
		field.Bool("status").Default(true).Comment("状态"),
		field.Int("sequence").Comment("序号"),
	}
}

// Edges of the ArticleCategory.
func (GoodArticleCategory) Edges() []ent.Edge {
	return nil
}

// Indexes of the ArticleCategory.
func (GoodArticleCategory) Indexes() []ent.Index {
	return []ent.Index{}
}

func (GoodArticleCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "good_article_category"},
	}
}

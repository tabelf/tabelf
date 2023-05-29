package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GoodArticleHot holds the schema definition for the Team entity.
type GoodArticleHot struct {
	common.BaseSchema
}

// Fields fields.
func (GoodArticleHot) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("article_uid").Comment("好文uid"),
		field.Bool("has_expired").Default(false).Comment("是否过期"),
	}
}

// Edges of the ArticleCategory.
func (GoodArticleHot) Edges() []ent.Edge {
	return nil
}

// Indexes of the ArticleCategory.
func (GoodArticleHot) Indexes() []ent.Index {
	return []ent.Index{}
}

func (GoodArticleHot) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "good_article_hot"},
	}
}

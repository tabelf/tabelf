package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// GoodArticleMeta holds the schema definition for the Team entity.
type GoodArticleMeta struct {
	common.BaseSchema
}

// Fields fields.
func (GoodArticleMeta) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("article_uid").Comment("好文uid"),
		field.String("user_uid").Comment("用户uid"),
		field.Bool("has_star").Default(false).Comment("是否收藏"),
		field.Bool("has_view").Default(false).Comment("是否查看"),
		field.Bool("has_used").Default(false).Comment("是否使用"),
	}
}

// Edges of the Account.
func (GoodArticleMeta) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (GoodArticleMeta) Indexes() []ent.Index {
	return []ent.Index{}
}

func (GoodArticleMeta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "good_article_meta"},
	}
}

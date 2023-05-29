package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UrlCrawl 允许收录的地址.
type UrlCrawl struct {
	common.BaseSchema
}

// Fields fields.
func (UrlCrawl) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("url").Comment("收录的url"),
		field.String("community_uid").Optional().Comment("分享社区uid"),
		field.Int("count").Default(0).Comment("收录次数"),
		field.Bool("has_complete").Default(false).Comment("是否收录完成"),
	}
}

// Edges of the Account.
func (UrlCrawl) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (UrlCrawl) Indexes() []ent.Index {
	return []ent.Index{}
}

func (UrlCrawl) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "url_crawl"},
	}
}

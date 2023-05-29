package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CommunityMeta holds the schema definition for the Team entity.
type CommunityMeta struct {
	common.BaseSchema
}

// Fields fields.
func (CommunityMeta) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("community_uid").Comment("社区uid"),
		field.String("user_uid").Comment("用户uid"),
		field.Bool("has_praise").Default(false).Comment("是否点赞"),
		field.Bool("has_star").Default(false).Comment("是否收藏"),
		field.Bool("has_view").Default(false).Comment("是否查看"),
		field.Bool("has_used").Default(false).Comment("是否使用"),
	}
}

// Edges of the Account.
func (CommunityMeta) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (CommunityMeta) Indexes() []ent.Index {
	return []ent.Index{}
}

func (CommunityMeta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "community_meta"},
	}
}

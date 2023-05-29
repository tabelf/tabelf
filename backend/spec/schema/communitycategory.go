package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// CommunityCategory holds the schema definition for the Team entity.
type CommunityCategory struct {
	common.BaseSchema
}

// Fields fields.
func (CommunityCategory) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Comment("分类名称"),
		field.Bool("status").Default(true).Comment("状态"),
		field.Int("sequence").Comment("序号"),
	}
}

// Edges of the Account.
func (CommunityCategory) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (CommunityCategory) Indexes() []ent.Index {
	return []ent.Index{}
}

func (CommunityCategory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "community_category"},
	}
}

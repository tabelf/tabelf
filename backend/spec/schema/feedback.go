package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Feedback holds the schema definition for the Feedback entity.
type Feedback struct {
	common.BaseSchema
}

// Fields fields.
func (Feedback) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("user_uid").Comment("用户uid"),
		field.String("category").Comment("问题分类"),
		field.String("description").Optional().Comment("问题描述"),
	}
}

// Edges of the Account.
func (Feedback) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Feedback) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
	}
}

func (Feedback) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "feedback"},
	}
}

package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Focus holds the schema definition for the Focus entity.
type Focus struct {
	common.BaseSchema
}

// Fields fields.
func (Focus) Fields() []ent.Field {
	return []ent.Field{
		field.String("follower_uid").Comment("关注者uid"),
		field.String("followee_uid").Comment("被关注者uid"),
		field.Bool("status").Default(false).Comment("关注状态"),
	}
}

// Edges of the Account.
func (Focus) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Focus) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("follower_uid"),
		index.Fields("followee_uid"),
	}
}

func (Focus) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "focus"},
	}
}

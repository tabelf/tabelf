package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TeamGroup holds the schema definition for the TeamGroup entity.
type TeamGroup struct {
	common.BaseSchema
}

// Fields fields.
func (TeamGroup) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("user_uid").Comment("用户uid"),
		field.String("team_uid").Comment("团队uid"),
	}
}

// Edges of the Account.
func (TeamGroup) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (TeamGroup) Indexes() []ent.Index {
	return []ent.Index{}
}

func (TeamGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "team_group"},
	}
}

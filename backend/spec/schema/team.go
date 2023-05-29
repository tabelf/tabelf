package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	common.BaseSchema
}

// Fields fields.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Comment("团队名称"),
		field.Time("expired_at").SchemaType(DatetimeSchema).Comment("有效期"),
	}
}

// Edges of the Account.
func (Team) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Team) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Team) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "team"},
	}
}

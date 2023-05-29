package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"tabelf/backend/common"
)

// Admin holds the schema definition for the Admin entity.
type Admin struct {
	common.BaseSchema
}

// Fields fields.
func (Admin) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_uid").Comment("管理员uid"),
	}
}

// Edges of the Account.
func (Admin) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Admin) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Admin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "admin"},
	}
}

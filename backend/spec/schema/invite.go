package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Invite holds the schema definition for the Invite entity.
type Invite struct {
	common.BaseSchema
}

// Fields fields.
func (Invite) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("referral_uid").Comment("邀请人uid"),
		field.String("referee_uid").Comment("注册者uid"),
	}
}

// Edges of the Invite.
func (Invite) Edges() []ent.Edge {
	return nil
}

// Indexes of the Invite.
func (Invite) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("referral_uid"),
	}
}

func (Invite) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "invite"},
	}
}

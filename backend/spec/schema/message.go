package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	common.BaseSchema
}

// Fields fields.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("user_uid").Comment("用户uid"),
		field.String("promoter_uid").Comment("发起人"),
		field.String("category").Comment("消息分类"),
		field.String("description").Comment("消息描述"),
		field.Bool("has_read").Default(false).Comment("是否已读"),
	}
}

// Edges of the Account.
func (Message) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
	}
}

func (Message) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "message"},
	}
}

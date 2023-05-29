package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PayOrderFeedback holds the schema definition for the PayOrderFeedback entity.
type PayOrderFeedback struct {
	common.BaseSchema
}

// Fields fields.
func (PayOrderFeedback) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("user_uid").Comment("用户uid"),
		field.String("order_number").Optional().Comment("订单号"),
		field.String("category").Comment("问题分类"),
		field.String("description").Optional().Comment("问题描述"),
	}
}

// Edges of the Account.
func (PayOrderFeedback) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (PayOrderFeedback) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
	}
}

func (PayOrderFeedback) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pay_order_feedback"},
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"tabelf/backend/common"
)

// Notice 系统发布公告.
type Notice struct {
	common.BaseSchema
}

// Fields fields.
func (Notice) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Comment("消息内容"),
		field.Bool("process").Default(false).Comment("进度"),
	}
}

// Edges of the Account.
func (Notice) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Notice) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Notice) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "notice"},
	}
}

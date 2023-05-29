package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"tabelf/backend/common"
)

// Recharge 充值消息.
type Recharge struct {
	common.BaseSchema
}

// Fields fields.
func (Recharge) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("title").Comment("标题"),
		field.String("origin_amount").SchemaType(DecimalSchema).Comment("原价"),
		field.String("amount").SchemaType(DecimalSchema).Comment("折扣价"),
		field.JSON("descriptions", []string{}).Comment("描述列表"),
		field.String("theme_color").Optional().Comment("开通按钮的主题色"),
		field.Int("year").Comment("开通增长年数限制"),
		field.Int("month").Comment("开通增长月数限制"),
		field.Bool("default").Default(false).Comment("是否作为默认金额"),
	}
}

// Edges of the Recharge.
func (Recharge) Edges() []ent.Edge {
	return nil
}

// Indexes of the Recharge.
func (Recharge) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Recharge) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "recharge"},
	}
}

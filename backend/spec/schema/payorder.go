package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"tabelf/backend/common"
)

// PayOrder 支付订单消息.
type PayOrder struct {
	common.BaseSchema
}

type RechargeRecord struct {
	UID          string   `json:"uid"`
	Title        string   `json:"title"`
	OriginAmount string   `json:"origin_amount"`
	Amount       string   `json:"amount"`
	Descriptions []string `json:"descriptions"`
	ThemeColor   string   `json:"theme_color"`
	Year         int      `json:"year"`
	Month        int      `json:"month"`
}

// Fields fields.
func (PayOrder) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("order_number").Comment("订单号"),
		field.String("order_type").Comment("订单类型"),
		field.String("user_uid").Comment("用户uid"),
		field.String("pay_method").Comment("支付方式"),
		field.String("payment_amount").SchemaType(DecimalSchema).Comment("实际支付金额"),
		field.String("total_price").SchemaType(DecimalSchema).Comment("总金额"),
		field.JSON("recharge_record", RechargeRecord{}).Comment("充值记录"),
		field.String("status").Validate(MaxRuneCount(LenTwenty)).Comment("订单状态"),
		field.String("transaction_number").Optional().Comment("YunGouOS 系统交易单号"),
		field.String("thirdparty_number").Optional().Comment("第三方支付单号"),
		field.String("mch_id").Optional().Comment("YunGouOS 商户号"),
		field.String("open_id").Optional().Comment("用户openID"),
		field.Time("member_expired").Optional().SchemaType(DatetimeSchema).Comment("会员过期时间"),
		field.String("cancel_event").Optional().Validate(MaxRuneCount(LenTwenty)).Comment("事件"),
	}
}

// Edges of the Recharge.
func (PayOrder) Edges() []ent.Edge {
	return nil
}

// Indexes of the Recharge.
func (PayOrder) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
		index.Fields("order_number"),
	}
}

func (PayOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pay_order"},
	}
}

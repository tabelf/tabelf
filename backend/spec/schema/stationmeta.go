package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// StationMeta holds the schema definition for the Team entity.
type StationMeta struct {
	common.BaseSchema
}

// Fields fields.
func (StationMeta) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("station_uid").Comment("推荐uid"),
		field.String("user_uid").Comment("用户uid"),
		field.Bool("has_praise").Default(false).Comment("是否点赞"),
		field.Bool("has_star").Default(false).Comment("是否收藏"),
		field.Bool("has_view").Default(false).Comment("是否查看"),
	}
}

// Edges of the Account.
func (StationMeta) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (StationMeta) Indexes() []ent.Index {
	return []ent.Index{}
}

func (StationMeta) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "station_meta"},
	}
}

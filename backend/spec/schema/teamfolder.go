package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// TeamFolder 团队文件.
type TeamFolder struct {
	common.BaseSchema
}

// Fields fields.
func (TeamFolder) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("team_uid").Comment("团队uid"),
		field.String("file_name").Optional().Comment("文件名称"),
	}
}

// Edges of the Account.
func (TeamFolder) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (TeamFolder) Indexes() []ent.Index {
	return []ent.Index{}
}

func (TeamFolder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "team_folder"},
	}
}

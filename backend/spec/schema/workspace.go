package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Workspace holds the schema definition for the Workspace entity.
type Workspace struct {
	common.BaseSchema
}

// Fields fields.
func (Workspace) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Comment("工作空间名称"),
		field.Int("type").Default(0).Comment("类型: 0 为个人, 1 为团队"),
		field.String("user_uid").Comment("用户uid"),
		field.String("personal_folder_uid").Optional().Comment("个人文件uid"),
		field.String("team_folder_uid").Optional().Comment("团队文件uid"),
		field.Bool("is_open").Default(false).Comment("是否打开, 0 折叠, 1 打开"),
	}
}

// Edges of the Account.
func (Workspace) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Workspace) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
		index.Fields("personal_folder_uid"),
	}
}

func (Workspace) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "workspace"},
	}
}

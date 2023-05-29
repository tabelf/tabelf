package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// PersonalFolder 个人文件.
type PersonalFolder struct {
	common.BaseSchema
}

// Fields fields.
func (PersonalFolder) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("user_uid").Comment("用户uid"),
		field.String("folder_name").SchemaType(VarcharFolderName).Comment("文件名称"),
		field.String("folder_number").MaxLen(LenFolderNumber).MinLen(LenFolderNumber).Comment("文件编号"),
		field.Bool("has_open").Default(false).Comment("是否开放到社区"),
	}
}

// Edges of the Account.
func (PersonalFolder) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (PersonalFolder) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_uid"),
		index.Fields("folder_number"),
		index.Fields("folder_number", "user_uid"),
	}
}

func (PersonalFolder) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "personal_folder"},
	}
}

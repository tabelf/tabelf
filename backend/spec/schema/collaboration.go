package schema

import (
	"entgo.io/ent/schema/index"
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Collaboration holds the schema definition for the Collaboration entity.
type Collaboration struct {
	common.BaseSchema
}

// Fields fields.
func (Collaboration) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("shard_uid").Comment("分享链接uid"),
		field.String("folder_uid").Comment("文件uid"),
		field.String("user_uid").Comment("用户uid"),
		field.String("authority").Comment("权限: 0 只读权限, 1 编辑权限"),
		field.String("folder_number").MaxLen(LenFolderNumber).MinLen(LenFolderNumber).Comment("文件编号"),
	}
}

// Edges of the Account.
func (Collaboration) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Collaboration) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("shard_uid"),
		index.Fields("folder_uid"),
		index.Fields("folder_number"),
	}
}

func (Collaboration) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "collaboration"},
	}
}

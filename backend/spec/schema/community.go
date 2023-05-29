package schema

import (
	"tabelf/backend/common"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Community holds the schema definition for the Team entity.
type Community struct {
	common.BaseSchema
}

// Fields fields.
func (Community) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("title").Comment("标题"),
		field.Text("description").Optional().Comment("描述"),
		field.String("image").Comment("图片"),
		field.JSON("tags", []string{}).Optional().Comment("标签列表"),
		field.Int("praise").Default(0).Comment("点赞量"),
		field.Int("star").Default(0).Comment("收藏量"),
		field.Int("view").Default(0).Comment("查看量"),
		field.Int("used").Default(0).Comment("使用量"),
		field.String("folder_uid").Comment("文件uid"),
		field.String("user_uid").Comment("用户uid"),
		field.String("status").Default("0").Comment("状态, -1 审核失败, 0 待审核, 1 审核通过"),
		field.String("remark").Optional().Comment("备注"),
		field.String("category_uid").Comment("分类uid"),
	}
}

// Edges of the Account.
func (Community) Edges() []ent.Edge {
	return nil
}

// Indexes of the Account.
func (Community) Indexes() []ent.Index {
	return []ent.Index{}
}

func (Community) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "community"},
	}
}

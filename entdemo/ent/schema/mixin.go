package schema

import (
	"entdemo/entrest"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/gogf/gf/util/guid"
)

// BaseMixin implements the ent.Mixin for sharing
type BaseMixin struct {
	mixin.Schema
}

// Fields of the BaseMixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			DefaultFunc(func() string {
				return guid.S()
			}).
			Comment("编号"),
		field.Time("createdAt").
			Immutable().
			Default(time.Now).
			Comment("创建时间").
			Annotations(
				entrest.Field{
					Creatable: false,
					Updatable: false,
				},
			),
		field.Time("updatedAt").
			Default(time.Now).
			UpdateDefault(time.Now).
			Comment("更新时间").
			Annotations(
				entrest.Field{
					Creatable: false,
					Updatable: false,
				},
			),
		field.Time("deletedAt").
			Optional().
			Nillable().
			Comment("删除时间").
			Annotations(
				entrest.Field{
					Creatable: false,
					Updatable: false,
				},
			),
		field.String("createdBy").
			Optional().
			Comment("创建者").
			Annotations(
				entrest.Field{
					Creatable: false,
					Updatable: false,
				},
			),
		field.String("updatedBy").
			Optional().
			Comment("修改者").
			Annotations(
				entrest.Field{
					Creatable: false,
					Updatable: false,
				},
			),
		field.String("remark").
			Optional().
			Comment("备注").
			Annotations(
				entrest.Field{
					Validate:  "required#请输入备注",
					Creatable: true,
					Updatable: true,
				},
			),
	}
}

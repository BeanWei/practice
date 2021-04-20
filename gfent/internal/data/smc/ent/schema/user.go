package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/BeanWei/gfent/internal/data/entx"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			Immutable().
			Unique().
			Comment("用户名"),
		field.String("password").
			Sensitive().
			Comment("密码"),
		field.String("phone").
			Nillable().
			Optional().
			Unique().
			Comment("手机号"),
		field.String("email").
			Nillable().
			Optional().
			Unique().
			Comment("邮箱"),
		field.String("avatar").
			Optional().
			Comment("头像"),
		field.Int("gender").
			Default(0).
			Comment("性别"),
		field.Bool("disabled").
			Default(false).
			Comment("是否禁用"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

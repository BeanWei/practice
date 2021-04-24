package schema

import (
	"entdemo/entrest"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the User
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.API{
			Prefix: "/api/v1",
		},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("用户名字").
			Annotations(
				entrest.Field{
					Validate:   "required#请输入用户的名字",
					Creatable:  true,
					Updatable:  false,
					Searchable: true,
				},
			),
		field.String("phone").
			Nillable().
			Optional().
			Comment("联系电话").
			Annotations(
				entrest.Field{
					Creatable:  true,
					Updatable:  true,
					Searchable: true,
				},
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type).
			Annotations(
				entrest.Edge{
					GetWith:   true,
					QueryWith: false,
				},
			),
	}
}

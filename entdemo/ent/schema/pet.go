package schema

import (
	"entdemo/entrest"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Mixin of the Pet
func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Annotations of the Pet
func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entrest.API{
			Prefix: "/api/v1",
			Middlewares: []*entrest.MiddlewareConfig{
				{
					PkgPath: "entdemo/middleware",
					Code:    "middleware.AllMiddleware",
				},
				{
					PkgPath: "entdemo/middleware",
					ReqType: entrest.ListRequest,
					Code:    "middleware.ListMiddleware",
				},
				{
					PkgPath: "entdemo/middleware",
					ReqType: entrest.GetRequest,
					Code:    "middleware.GetMiddleware",
				},
				{
					PkgPath: "entdemo/middleware",
					ReqType: entrest.CreateRequest,
					Code:    "middleware.CreateMiddleware",
				},
				{
					PkgPath: "entdemo/middleware",
					ReqType: entrest.UpdateRequest,
					Code:    "middleware.UpdateMiddleware",
				},
				{
					PkgPath: "entdemo/middleware",
					ReqType: entrest.DeleteRequest,
					Code:    "middleware.DeleteMiddleware",
				},
			},
		},
	}
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("宠物名字").
			Annotations(
				entrest.Field{
					Validate:   "required#请输入宠物的名字",
					Creatable:  true,
					Updatable:  true,
					Searchable: true,
				},
			),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}

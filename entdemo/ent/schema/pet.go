package schema

import (
	"entdemo/entrest"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(
				entrest.Annotation{
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

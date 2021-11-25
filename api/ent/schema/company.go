package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("business_number"),
		field.String("address").Optional().Nillable(),
		field.String("description").Optional().Nillable(),
		field.String("profile_image").Optional().Nillable(),
		field.String("hompage").Optional().Nillable(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("company").
			Required().
			Unique(),
	}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

var kinds = []string{"company", "student"}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("password"),
		field.Enum("kind").Values(kinds...),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Club holds the schema definition for the Club entity.
type Club struct {
	ent.Schema
}

// Fields of the Club.
func (Club) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("organization"),
		field.String("description"),
		field.String("profile_image").Optional().Nillable(),
		field.String("profile_link").Optional().Nillable(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
	}
}

// Edges of the Club.
func (Club) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("leader", Student.Type).
			Ref("club").
			Required().
			Unique(),
		edge.To("club_member", ClubMember.Type),
	}
}

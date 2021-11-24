package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ClubMember holds the schema definition for the ClubMember entity.
type ClubMember struct {
	ent.Schema
}

// Fields of the ClubMember.
func (ClubMember) Fields() []ent.Field {
	return []ent.Field{
		field.Int("club_id"),
		field.Int("student_id"),
	}
}

// Edges of the ClubMember.
func (ClubMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).
			Ref("club_member").
			Field("student_id").
			Unique().
			Required(),
		edge.From("club", Club.Type).
			Ref("club_member").
			Field("club_id").
			Unique().
			Required(),
	}
}

func (ClubMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("club_id", "student_id").
			Unique(),
	}
}

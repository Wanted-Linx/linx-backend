package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ProjectClub holds the schema definition for the ProjectClub entity.
type ProjectClub struct {
	ent.Schema
}

// Fields of the ProjectClub.
func (ProjectClub) Fields() []ent.Field {
	return []ent.Field{
		field.Int("project_id"),
		field.Int("club_id"),
		field.String("start_date"),
	}
}

// Edges of the ProjectClub.
func (ProjectClub) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("club", Club.Type).
			Ref("project_club").
			Field("club_id").
			Required().
			Unique(),
		edge.From("project", Project.Type).
			Ref("project_club").
			Field("project_id").
			Required().
			Unique(),
	}
}

func (ProjectClub) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("project_id", "club_id").
			Unique(),
	}
}

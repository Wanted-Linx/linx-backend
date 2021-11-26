package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProjectLogParticipant holds the schema definition for the ProjectLogParticipant entity.
type ProjectLogParticipant struct {
	ent.Schema
}

// Fields of the ProjectLogParticipant.
func (ProjectLogParticipant) Fields() []ent.Field {
	return []ent.Field{
		// field.Int("student_id") PK,FK?
		field.String("name"),
	}
}

// Edges of the ProjectLogParticipant.
func (ProjectLogParticipant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project_log", ProjectLog.Type).
			Ref("project_log_participant").
			Required().
			Unique(),
	}
}

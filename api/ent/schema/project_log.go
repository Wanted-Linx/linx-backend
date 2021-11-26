package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProjectLog holds the schema definition for the ProjectLog entity.
type ProjectLog struct {
	ent.Schema
}

// Fields of the ProjectLog.
func (ProjectLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("author"),
		field.Text("content"),
		field.String("start_date"),
		field.String("end_date"),
		field.Time("created_at").Optional(),
	}
}

// Edges of the ProjectLog.
func (ProjectLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project", Project.Type).
			Ref("project_log").
			Required().
			Unique(),
		edge.From("project_club", ProjectClub.Type).
			Ref("project_log").
			Required().
			Unique(),
		edge.To("project_log_participant", ProjectLogParticipant.Type),
		edge.To("project_log_feedback", ProjectLogFeedback.Type),
	}
}

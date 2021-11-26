package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProjectLogFeedback holds the schema definition for the ProjectLogFeedback entity.
type ProjectLogFeedback struct {
	ent.Schema
}

// Fields of the ProjectLogFeedback.
func (ProjectLogFeedback) Fields() []ent.Field {
	return []ent.Field{
		field.String("author"),
		field.Text("content"),
	}
}

// Edges of the ProjectLogFeedback.
func (ProjectLogFeedback) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("project_log", ProjectLog.Type).
			Ref("project_log_feedback").
			Required().
			Unique(),
	}
}

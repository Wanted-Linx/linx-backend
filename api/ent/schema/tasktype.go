package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TaskType holds the schema definition for the TaskType entity.
type TaskType struct {
	ent.Schema
}

// Fields of the TaskType.
func (TaskType) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
	}
}

// Edges of the TaskType.
func (TaskType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("student", Student.Type).
			Ref("task_type").Unique(),
		edge.From("club", Club.Type).
			Ref("task_type").Unique(),
		edge.From("company", Company.Type).
			Ref("task_type").Unique(),
		edge.From("project", Project.Type).
			Ref("task_type").Unique(),
	}
}

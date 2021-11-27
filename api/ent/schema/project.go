package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Text("content"),
		field.String("start_date"),
		field.String("end_date"),
		field.String("applying_start_date"),
		field.String("applying_end_date"),
		field.String("qualification"),
		field.String("profile_image").Optional().Nillable(),
		field.Time("created_at").Default(func() time.Time {
			return time.Now()
		}),
		field.Int("sponsor_fee"),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("company", Company.Type).
			Ref("project").
			Required().
			Unique(),
		edge.From("club", Club.Type).
			Ref("project").
			Unique(),
		// Club Required 뺌(프로젝트 최초 생성할 때는 없으니까...)
		edge.To("project_club", ProjectClub.Type),
		edge.To("project_log", ProjectLog.Type),
	}
}

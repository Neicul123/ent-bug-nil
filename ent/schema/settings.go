package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Int("list_id"),
		field.String("test"),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("list", List.Type).Unique().Required().Field("list_id"),
		edge.To("user", User.Type).Unique().Required().Field("user_id"),
	}
}

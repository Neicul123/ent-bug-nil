package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// List holds the schema definition for the List entity.
type List struct {
	ent.Schema
}

// Fields of the List.
func (List) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
	}
}

// Edges of the List.
func (List) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("studying", User.Type).Ref("studies").Through("settings", Settings.Type),
	}
}

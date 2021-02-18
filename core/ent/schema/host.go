package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname").Unique(),
		field.String("interface"),
		field.Int("lastSeen").Default(0),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bots", Bot.Type),
		edge.To("actions", Action.Type),
		edge.From("member", Group.Type).
			Ref("members"),
	}
}

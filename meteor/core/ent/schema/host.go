package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.String("hostname"),
		field.String("interface"),
		field.Int("lastSeen").Positive(),
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

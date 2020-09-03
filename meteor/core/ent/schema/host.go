package schema

import (
	"github.com/facebook/ent"
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
	return nil
}

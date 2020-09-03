package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.String("mode"),
		field.String("args"),
		field.Bool("queued"),
		field.Bool("responded"),
		field.String("response"),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return nil
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Action holds the schema definition for the Action entity.
type Action struct {
	ent.Schema
}

// Fields of the Action.
func (Action) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.String("mode"),
		field.String("args"),
		field.Bool("queued").Default(false),
		field.Bool("responded").Default(false),
		field.String("result").Default("N/A"),
	}
}

// Edges of the Action.
func (Action) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("targeting", Host.Type).
			Ref("actions").Unique(),
	}
}

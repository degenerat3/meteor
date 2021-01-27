package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Bot holds the schema definition for the Bot entity.
type Bot struct {
	ent.Schema
}

// Fields of the Bot.
func (Bot) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid").Unique(),
		field.Int("interval"),
		field.Int("delta"),
		field.Int("lastSeen").Default(0),
	}
}

// Edges of the Bot.
func (Bot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("infecting", Host.Type).
			Ref("bots").Unique(),
	}
}

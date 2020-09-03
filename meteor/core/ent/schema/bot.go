package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Bot holds the schema definition for the Bot entity.
type Bot struct {
	ent.Schema
}

// Fields of the Bot.
func (Bot) Fields() []ent.Field {
	return []ent.Field{
		field.String("uuid"),
		field.Int("interval").Positive(),
		field.Int("delta").Positive(),
		field.Int("lastSeen").Positive(),
	}
}

// Edges of the Bot.
func (Bot) Edges() []ent.Edge {
	return nil
}

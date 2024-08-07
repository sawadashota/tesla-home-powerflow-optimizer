package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChargeCommandHistory holds the schema definition for the ChargeCommandHistory entity.
type ChargeCommandHistory struct {
	ent.Schema
}

// Fields of the ChargeCommandHistory.
func (ChargeCommandHistory) Fields() []ent.Field {
	return []ent.Field{
		field.String("vin"),
		field.String("operation"),
		field.Int("amps"),
		field.Time("timestamp"),
	}
}

// Edges of the ChargeCommandHistory.
func (ChargeCommandHistory) Edges() []ent.Edge {
	return nil
}

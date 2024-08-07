package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChargeSetting holds the schema definition for the ChargeSetting entity.
type ChargeSetting struct {
	ent.Schema
}

// Fields of the ChargeSetting.
func (ChargeSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("enabled").Default(false),
		field.Int("charge_start_threshold"),
		field.Int("power_usage_increase_threshold"),
		field.Int("power_usage_decrease_threshold"),
		field.Int("update_interval"),
	}
}

// Edges of the ChargeSetting.
func (ChargeSetting) Edges() []ent.Edge {
	return nil
}

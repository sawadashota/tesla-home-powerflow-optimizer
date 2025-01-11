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
		field.Int("min_charge_threshold").Min(0).Max(100).Default(0),
		field.String("min_charge_time_range_start").Default("21:00"),
		field.String("min_charge_time_range_end").Default("06:00"),
		field.Int("min_charge_amperage").Min(5).Max(16).Default(16),
	}
}

// Edges of the ChargeSetting.
func (ChargeSetting) Edges() []ent.Edge {
	return nil
}

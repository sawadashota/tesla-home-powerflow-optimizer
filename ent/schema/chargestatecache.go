package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ChargeStateCache holds the schema definition for the ChargeStateCache entity.
type ChargeStateCache struct {
	ent.Schema
}

// Fields of the ChargeStateCache.
func (ChargeStateCache) Fields() []ent.Field {
	return []ent.Field{
		field.String("vin").Unique(),
		field.Int("battery_level"),
		field.Float32("battery_range"),
		field.Int("charge_amps"),
		field.Int("charge_current_request"),
		field.Int("charge_current_request_max"),
		field.Bool("charge_enable_request"),
		field.Int("charge_limit_soc"),
		field.Bool("charge_port_door_open"),
		field.String("charge_port_latch"),
		field.Int("charger_actual_current"),
		field.Int("charger_voltage"),
		field.String("charging_state"),
		field.Int("minutes_to_full_charge"),
		field.Time("timestamp"),
		field.Int("usable_battery_level"),
	}
}

// Edges of the ChargeStateCache.
func (ChargeStateCache) Edges() []ent.Edge {
	return nil
}

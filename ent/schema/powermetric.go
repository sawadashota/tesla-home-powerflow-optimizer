package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// PowerMetric holds the schema definition for the PowerMetric entity.
type PowerMetric struct {
	ent.Schema
}

// Fields of the PowerMetric.
func (PowerMetric) Fields() []ent.Field {
	return []ent.Field{
		field.Int("surplus_watt"),
		field.Time("timestamp"),
	}
}

// Edges of the PowerMetric.
func (PowerMetric) Edges() []ent.Edge {
	return nil
}

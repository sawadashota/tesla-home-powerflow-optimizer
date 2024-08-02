package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Grant holds the schema definition for the Grant entity.
type Grant struct {
	ent.Schema
}

// Fields of the Grant.
func (Grant) Fields() []ent.Field {
	return []ent.Field{
		field.String("subject").NotEmpty().Unique(),
		field.String("access_token").NotEmpty().Unique(),
		field.String("refresh_token").NotEmpty(),
		field.String("scope").NotEmpty(),
		field.Time("expiry"),
	}
}

// Edges of the Grant.
func (Grant) Edges() []ent.Edge {
	return nil
}

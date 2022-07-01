package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Time("created_at"),
		field.Int("created_by"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}

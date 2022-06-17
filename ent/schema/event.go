package schema

import {
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/schema/index"
}


type Event struct {
	ent.schema
}

func (Event) Fields() []ent.Field {
	return []ent.Field {
		field.String("name"),
		field.String("description"),
		field.Time("start_date"),
		field.Time("end_date"),
		field.Time("created_at").
		Default(time.Now).
		Immutable(),
		field.int("created_by")
	}
}

func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("config", Config.Type)
		.Unique(),
	}
}
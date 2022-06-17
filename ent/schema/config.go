package schema


import {
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/schema/index"

	"github.com/jackc/pgtype"
}

type Config struct {
	ent.Schema
}

func (Config) Fields []ent.Field {
	return []ent.Field {
		field.int("days"),
		field.Other("start", &pgtype.Time).
		schemaType(map[string]string {
			dialect.Postgres: "time"
		}),
		field.Other("end",  &pgtype.Time).
		schemaType(map[string]string {
			dialect.Postgres: "time"
		}),
	}
}
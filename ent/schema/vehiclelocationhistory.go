package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// VehicleLocationHistory holds the schema definition for the VehicleLocationHistory entity.
type VehicleLocationHistory struct {
	ent.Schema
}

func (VehicleLocationHistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "vehicle_location_histories"},
	}
}

// Fields of the VehicleLocationHistory.
func (VehicleLocationHistory) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.String("city"),

		// Foreign key UUID has no default!
		field.UUID("ride_id", uuid.UUID{}),

		field.Time("timestamp").
			// Do not convert to session timezone
			SchemaType(map[string]string{
				dialect.Postgres: "TIMESTAMP",
			}),

		field.Float("lat").Optional().
			// Store as specifc float type
			SchemaType(map[string]string{
				dialect.Postgres: "FLOAT8",
			}),

		field.Float("long").Optional().
			// Store as specifc float type
			SchemaType(map[string]string{
				dialect.Postgres: "FLOAT8",
			}),
	}
}

// Edges of the VehicleLocationHistory.
func (VehicleLocationHistory) Edges() []ent.Edge {
	return []ent.Edge{

		// A VehicleLocationHistory has one ride
		edge.From("rides", Ride.Type).
			Ref("vehicle_location_histories").
			Field("ride_id").
			Unique().
			Required(),
	}
}

// Indexes of the Vehicle.
func (VehicleLocationHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("city", "ride_id", "timestamp"),
	}
}

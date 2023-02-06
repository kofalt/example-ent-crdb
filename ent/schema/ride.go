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

// Ride holds the schema definition for the Ride entity.
type Ride struct {
	ent.Schema
}

func (Ride) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "rides"},
	}
}

// Fields of the Ride.
func (Ride) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.String("city"),
		field.String("vehicle_city").Optional(),

		// Foreign key optional UUID has no default!
		field.UUID("rider_id", uuid.UUID{}).Optional(),
		field.UUID("vehicle_id", uuid.UUID{}).Optional(),

		field.String("start_address").Optional(),
		field.String("end_address").Optional(),

		field.Time("start_time").Optional().
			// Do not convert to session timezone
			SchemaType(map[string]string{
				dialect.Postgres: "TIMESTAMP",
			}),

		field.Time("end_time").Optional().
			// Do not convert to session timezone
			SchemaType(map[string]string{
				dialect.Postgres: "TIMESTAMP",
			}),

		field.Float("revenue").Optional().
			// Store as fixed-point
			SchemaType(map[string]string{
				dialect.Postgres: "decimal(10,2)",
			}),
	}
}

// Edges of the Ride.
func (Ride) Edges() []ent.Edge {
	return []ent.Edge{

		// A ride has zero or one user (rider)
		edge.To("user", User.Type).
			Field("rider_id").
			Unique(),

		// A ride has zero or one vehicle
		edge.To("vehicle", Vehicle.Type).
			Field("vehicle_id").
			Unique(),

		// A Ride has one VehicleLocationHistory
		edge.To("vehicle_location_histories", VehicleLocationHistory.Type),
	}

}

// Indexes of the Ride.
func (Ride) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("city", "id"),
		// index.Fields("city", "owner_id"),
	}
}

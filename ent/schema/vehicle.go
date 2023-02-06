package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// Vehicle holds the schema definition for the Vehicle entity.
type Vehicle struct {
	ent.Schema
}

// Fields of the Vehicle.
func (Vehicle) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.String("city"),
		field.String("type").Optional(),

		// Foreign key optional UUID has no default!
		field.UUID("owner_id", uuid.UUID{}).Optional(),

		field.Time("creation_time").Optional().
			// Do not convert to session timezone
			SchemaType(map[string]string{
				dialect.Postgres: "TIMESTAMP",
			}),

		field.String("status").Optional(),
		field.String("current_location").Optional(),

		field.JSON("ext", &map[string]interface{}{}).Optional(),
	}
}

// Edges of the Vehicle.
func (Vehicle) Edges() []ent.Edge {
	return []ent.Edge{

		// A vehicle has zero or one users
		edge.From("user", User.Type).
			Ref("vehicles").
			Field("owner_id").
			Unique(),

		// A vehicle has zero or more rides
		edge.From("rides", Ride.Type).Ref("vehicle"),
	}

}

// Indexes of the Vehicle.
func (Vehicle) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("city", "id"),
		index.Fields("city", "owner_id"),
	}
}

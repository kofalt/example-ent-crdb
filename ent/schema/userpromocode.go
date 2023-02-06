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

// UserPromoCode holds the schema definition for the UserPromoCode entity.
type UserPromoCode struct {
	ent.Schema
}

func (UserPromoCode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_promo_codes"},
	}
}

// Fields of the UserPromoCode.
func (UserPromoCode) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.String("city"),

		// Foreign key UUID has no default!
		field.UUID("user_id", uuid.UUID{}),

		field.String("code"),

		field.Time("timestamp").Optional().
			// Do not convert to session timezone
			SchemaType(map[string]string{
				dialect.Postgres: "TIMESTAMP",
			}),

		field.Int("usage_count").Optional(),
	}
}

// Edges of the VehicleLocationHistory.
func (UserPromoCode) Edges() []ent.Edge {
	return []ent.Edge{

		// A UserPromoCode has one ride
		edge.From("user", User.Type).
			Ref("user_promo_codes").
			Field("user_id").
			Unique().
			Required(),
	}
}

// Indexes of the Vehicle.
func (UserPromoCode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("city", "user_id", "code"),
	}
}

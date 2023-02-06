package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "users"},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.String("city"),
		field.String("name").Optional(),
		field.String("address").Optional(),
		field.String("credit_card").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{

		// A user has zero or more vehicles
		edge.To("vehicles", Vehicle.Type),

		// A user has zero or more rides
		edge.From("rides", Ride.Type).Ref("user"),

		// A user has zero or more promo codes
		edge.To("user_promo_codes", UserPromoCode.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("city", "id"),
	}
}

// Code generated by ent, DO NOT EDIT.

package vehicle

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kofalt/example-ent-crdb/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLTE(FieldID, id))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldCity, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldType, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldOwnerID, v))
}

// CreationTime applies equality check predicate on the "creation_time" field. It's identical to CreationTimeEQ.
func CreationTime(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldCreationTime, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldStatus, v))
}

// CurrentLocation applies equality check predicate on the "current_location" field. It's identical to CurrentLocationEQ.
func CurrentLocation(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldCurrentLocation, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasSuffix(FieldCity, v))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContainsFold(FieldCity, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContainsFold(FieldType, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...uuid.UUID) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDIsNil applies the IsNil predicate on the "owner_id" field.
func OwnerIDIsNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIsNull(FieldOwnerID))
}

// OwnerIDNotNil applies the NotNil predicate on the "owner_id" field.
func OwnerIDNotNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotNull(FieldOwnerID))
}

// CreationTimeEQ applies the EQ predicate on the "creation_time" field.
func CreationTimeEQ(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldCreationTime, v))
}

// CreationTimeNEQ applies the NEQ predicate on the "creation_time" field.
func CreationTimeNEQ(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldCreationTime, v))
}

// CreationTimeIn applies the In predicate on the "creation_time" field.
func CreationTimeIn(vs ...time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldCreationTime, vs...))
}

// CreationTimeNotIn applies the NotIn predicate on the "creation_time" field.
func CreationTimeNotIn(vs ...time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldCreationTime, vs...))
}

// CreationTimeGT applies the GT predicate on the "creation_time" field.
func CreationTimeGT(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGT(FieldCreationTime, v))
}

// CreationTimeGTE applies the GTE predicate on the "creation_time" field.
func CreationTimeGTE(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGTE(FieldCreationTime, v))
}

// CreationTimeLT applies the LT predicate on the "creation_time" field.
func CreationTimeLT(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLT(FieldCreationTime, v))
}

// CreationTimeLTE applies the LTE predicate on the "creation_time" field.
func CreationTimeLTE(v time.Time) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLTE(FieldCreationTime, v))
}

// CreationTimeIsNil applies the IsNil predicate on the "creation_time" field.
func CreationTimeIsNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIsNull(FieldCreationTime))
}

// CreationTimeNotNil applies the NotNil predicate on the "creation_time" field.
func CreationTimeNotNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotNull(FieldCreationTime))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotNull(FieldStatus))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContainsFold(FieldStatus, v))
}

// CurrentLocationEQ applies the EQ predicate on the "current_location" field.
func CurrentLocationEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEQ(FieldCurrentLocation, v))
}

// CurrentLocationNEQ applies the NEQ predicate on the "current_location" field.
func CurrentLocationNEQ(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNEQ(FieldCurrentLocation, v))
}

// CurrentLocationIn applies the In predicate on the "current_location" field.
func CurrentLocationIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIn(FieldCurrentLocation, vs...))
}

// CurrentLocationNotIn applies the NotIn predicate on the "current_location" field.
func CurrentLocationNotIn(vs ...string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotIn(FieldCurrentLocation, vs...))
}

// CurrentLocationGT applies the GT predicate on the "current_location" field.
func CurrentLocationGT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGT(FieldCurrentLocation, v))
}

// CurrentLocationGTE applies the GTE predicate on the "current_location" field.
func CurrentLocationGTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldGTE(FieldCurrentLocation, v))
}

// CurrentLocationLT applies the LT predicate on the "current_location" field.
func CurrentLocationLT(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLT(FieldCurrentLocation, v))
}

// CurrentLocationLTE applies the LTE predicate on the "current_location" field.
func CurrentLocationLTE(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldLTE(FieldCurrentLocation, v))
}

// CurrentLocationContains applies the Contains predicate on the "current_location" field.
func CurrentLocationContains(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContains(FieldCurrentLocation, v))
}

// CurrentLocationHasPrefix applies the HasPrefix predicate on the "current_location" field.
func CurrentLocationHasPrefix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasPrefix(FieldCurrentLocation, v))
}

// CurrentLocationHasSuffix applies the HasSuffix predicate on the "current_location" field.
func CurrentLocationHasSuffix(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldHasSuffix(FieldCurrentLocation, v))
}

// CurrentLocationIsNil applies the IsNil predicate on the "current_location" field.
func CurrentLocationIsNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIsNull(FieldCurrentLocation))
}

// CurrentLocationNotNil applies the NotNil predicate on the "current_location" field.
func CurrentLocationNotNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotNull(FieldCurrentLocation))
}

// CurrentLocationEqualFold applies the EqualFold predicate on the "current_location" field.
func CurrentLocationEqualFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldEqualFold(FieldCurrentLocation, v))
}

// CurrentLocationContainsFold applies the ContainsFold predicate on the "current_location" field.
func CurrentLocationContainsFold(v string) predicate.Vehicle {
	return predicate.Vehicle(sql.FieldContainsFold(FieldCurrentLocation, v))
}

// ExtIsNil applies the IsNil predicate on the "ext" field.
func ExtIsNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldIsNull(FieldExt))
}

// ExtNotNil applies the NotNil predicate on the "ext" field.
func ExtNotNil() predicate.Vehicle {
	return predicate.Vehicle(sql.FieldNotNull(FieldExt))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRides applies the HasEdge predicate on the "rides" edge.
func HasRides() predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RidesTable, RidesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRidesWith applies the HasEdge predicate on the "rides" edge with a given conditions (other predicates).
func HasRidesWith(preds ...predicate.Ride) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RidesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RidesTable, RidesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vehicle) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vehicle) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vehicle) predicate.Vehicle {
	return predicate.Vehicle(func(s *sql.Selector) {
		p(s.Not())
	})
}

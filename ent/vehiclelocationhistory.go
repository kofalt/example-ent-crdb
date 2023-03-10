// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/kofalt/example-ent-crdb/ent/ride"
	"github.com/kofalt/example-ent-crdb/ent/vehiclelocationhistory"
)

// VehicleLocationHistory is the model entity for the VehicleLocationHistory schema.
type VehicleLocationHistory struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// RideID holds the value of the "ride_id" field.
	RideID uuid.UUID `json:"ride_id,omitempty"`
	// Timestamp holds the value of the "timestamp" field.
	Timestamp time.Time `json:"timestamp,omitempty"`
	// Lat holds the value of the "lat" field.
	Lat float64 `json:"lat,omitempty"`
	// Long holds the value of the "long" field.
	Long float64 `json:"long,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VehicleLocationHistoryQuery when eager-loading is set.
	Edges VehicleLocationHistoryEdges `json:"edges"`
}

// VehicleLocationHistoryEdges holds the relations/edges for other nodes in the graph.
type VehicleLocationHistoryEdges struct {
	// Rides holds the value of the rides edge.
	Rides *Ride `json:"rides,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// RidesOrErr returns the Rides value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VehicleLocationHistoryEdges) RidesOrErr() (*Ride, error) {
	if e.loadedTypes[0] {
		if e.Rides == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: ride.Label}
		}
		return e.Rides, nil
	}
	return nil, &NotLoadedError{edge: "rides"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VehicleLocationHistory) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vehiclelocationhistory.FieldLat, vehiclelocationhistory.FieldLong:
			values[i] = new(sql.NullFloat64)
		case vehiclelocationhistory.FieldCity:
			values[i] = new(sql.NullString)
		case vehiclelocationhistory.FieldTimestamp:
			values[i] = new(sql.NullTime)
		case vehiclelocationhistory.FieldID, vehiclelocationhistory.FieldRideID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type VehicleLocationHistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VehicleLocationHistory fields.
func (vlh *VehicleLocationHistory) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vehiclelocationhistory.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				vlh.ID = *value
			}
		case vehiclelocationhistory.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				vlh.City = value.String
			}
		case vehiclelocationhistory.FieldRideID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ride_id", values[i])
			} else if value != nil {
				vlh.RideID = *value
			}
		case vehiclelocationhistory.FieldTimestamp:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field timestamp", values[i])
			} else if value.Valid {
				vlh.Timestamp = value.Time
			}
		case vehiclelocationhistory.FieldLat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field lat", values[i])
			} else if value.Valid {
				vlh.Lat = value.Float64
			}
		case vehiclelocationhistory.FieldLong:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field long", values[i])
			} else if value.Valid {
				vlh.Long = value.Float64
			}
		}
	}
	return nil
}

// QueryRides queries the "rides" edge of the VehicleLocationHistory entity.
func (vlh *VehicleLocationHistory) QueryRides() *RideQuery {
	return NewVehicleLocationHistoryClient(vlh.config).QueryRides(vlh)
}

// Update returns a builder for updating this VehicleLocationHistory.
// Note that you need to call VehicleLocationHistory.Unwrap() before calling this method if this VehicleLocationHistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (vlh *VehicleLocationHistory) Update() *VehicleLocationHistoryUpdateOne {
	return NewVehicleLocationHistoryClient(vlh.config).UpdateOne(vlh)
}

// Unwrap unwraps the VehicleLocationHistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vlh *VehicleLocationHistory) Unwrap() *VehicleLocationHistory {
	_tx, ok := vlh.config.driver.(*txDriver)
	if !ok {
		panic("ent: VehicleLocationHistory is not a transactional entity")
	}
	vlh.config.driver = _tx.drv
	return vlh
}

// String implements the fmt.Stringer.
func (vlh *VehicleLocationHistory) String() string {
	var builder strings.Builder
	builder.WriteString("VehicleLocationHistory(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vlh.ID))
	builder.WriteString("city=")
	builder.WriteString(vlh.City)
	builder.WriteString(", ")
	builder.WriteString("ride_id=")
	builder.WriteString(fmt.Sprintf("%v", vlh.RideID))
	builder.WriteString(", ")
	builder.WriteString("timestamp=")
	builder.WriteString(vlh.Timestamp.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("lat=")
	builder.WriteString(fmt.Sprintf("%v", vlh.Lat))
	builder.WriteString(", ")
	builder.WriteString("long=")
	builder.WriteString(fmt.Sprintf("%v", vlh.Long))
	builder.WriteByte(')')
	return builder.String()
}

// VehicleLocationHistories is a parsable slice of VehicleLocationHistory.
type VehicleLocationHistories []*VehicleLocationHistory

func (vlh VehicleLocationHistories) config(cfg config) {
	for _i := range vlh {
		vlh[_i].config = cfg
	}
}

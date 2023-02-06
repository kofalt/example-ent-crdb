package data

import (
	// "github.com/rs/zerolog/log"

	"github.com/google/uuid"
	"github.com/kofalt/example-ent-crdb/ent"
	"github.com/kofalt/example-ent-crdb/ent/user"
	. "github.com/kofalt/example-ent-crdb/util"
)

type Paginate struct {
	Offset int
	Limit  int
}

type UserSearch struct {
	*Paginate
	City string
}

// Retrieve paginated users with optional search
func (db *DB) GetUsers(s *UserSearch) ([]*ent.User, error) {
	q := db.User.Query().
		Order(ent.Desc("id")).
		Offset(s.Offset).Limit(s.Limit)

	if s.City != "" {
		q = q.Where(user.CityContains(s.City))
	}

	return q.All(BG)
}

// Retrieve a single user by ID
func (db *DB) GetUser(id uuid.UUID) (*ent.User, error) {
	return db.User.Query().
		Where(user.ID(id)).
		Only(BG)
}

// Move a user to a new address
func (db *DB) MoveUser(u *ent.User) (*ent.User, error) {
	return db.Debug().User.UpdateOne(u).
		SetAddress(u.Address).
		Save(BG)
}

// Traverse the graph to show the top ride(s) and their relations.
func (db *DB) TopRides() ([]*ent.Ride, error) {
	q := db.Ride.Query().
		Order(ent.Desc("revenue")).Limit(1).

		// Join users
		WithUser().

		// Join vehicle
		WithVehicle()

	return q.All(BG)
}

// q = q.WithRides(func(q *ent.RideQuery) {
// 	q.Limit(1)
// })

package data

import (
	"strconv"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/kofalt/example-ent-crdb/ent"
	. "github.com/kofalt/example-ent-crdb/util"
)

func (db *DB) CountTables(msg string) int {
	users, err1 := db.User.Query().Count(BG)
	vehicles, err2 := db.Vehicle.Query().Count(BG)
	rides, err3 := db.Ride.Query().Count(BG)
	Fatal(err1, err2, err3)

	log.Debug().
		Int("users", users).
		Int("vehicles", vehicles).
		Int("rides", rides).
		Msg(msg)

	return users + vehicles + rides
}

func (db *DB) CheckPreload() error {
	total := db.CountTables("Prexisting data")

	if total == 0 {
		Fatal(db.Preload())
		db.CountTables("Loaded data")
	}

	return nil
}

func (db *DB) Preload() error {

	bulkU := make([]*ent.UserCreate, len(MovrUsers))
	for i, x := range MovrUsers {
		bulkU[i] = db.User.Create().
			SetID(uuid.MustParse(x["id"])).
			SetCity(x["city"]).
			SetName(x["name"]).
			SetAddress(x["address"]).
			SetCreditCard(x["credit_card"])
	}

	bulkV := make([]*ent.VehicleCreate, len(MovrVehicles))
	for i, x := range MovrVehicles {

		t, err := time.Parse(time.RFC3339, x["creation_time"])
		Fatal(err)

		ext, err := DecodeStringMap([]byte(x["ext"]))
		Fatal(err)

		bulkV[i] = db.Vehicle.Create().
			SetID(uuid.MustParse(x["id"])).
			SetCity(x["city"]).
			SetType(x["type"]).
			SetUserID(uuid.MustParse(x["owner_id"])).
			SetCreationTime(t).
			SetStatus(x["status"]).
			SetCurrentLocation(x["current_location"]).
			SetExt(&ext)
	}

	bulkR := make([]*ent.RideCreate, len(MovrRides))
	for i, x := range MovrRides {

		t1, err1 := time.Parse(time.RFC3339, x["start_time"])
		t2, err2 := time.Parse(time.RFC3339, x["end_time"])
		rev, err3 := strconv.ParseFloat(x["revenue"], 64)
		Fatal(err1, err2, err3)

		bulkR[i] = db.Ride.Create().
			SetID(uuid.MustParse(x["id"])).
			SetCity(x["city"]).
			SetVehicleCity(x["vehicle_city"]).
			SetRiderID(uuid.MustParse(x["rider_id"])).
			SetVehicleID(uuid.MustParse(x["vehicle_id"])).
			SetStartAddress(x["start_address"]).
			SetEndAddress(x["end_address"]).
			SetStartTime(t1).
			SetEndTime(t2).
			SetRevenue(rev)
	}

	_, err1 := db.User.CreateBulk(bulkU...).Save(BG)
	_, err2 := db.Vehicle.CreateBulk(bulkV...).Save(BG)
	_, err3 := db.Ride.CreateBulk(bulkR...).Save(BG)
	Fatal(err1, err2, err3)

	return nil
}

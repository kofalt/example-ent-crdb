package data

import (
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"

	"github.com/kofalt/example-ent-crdb/ent"
	. "github.com/kofalt/example-ent-crdb/util"
)

type DB struct {
	*ent.Client
}

func Setup(s *DbConnectSettings) (*DB, error) {
	client, err := ent.Open("postgres", s.String())
	Fatal(err)

	db := &DB{
		client,
	}

	// Schema migration
	err = db.Schema.Create(BG, schema.WithAtlas(true))
	Fatal(err)

	// Preload data
	err = db.CheckPreload()
	Fatal(err)

	return db, nil
}

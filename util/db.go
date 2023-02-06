package util

import (
	"fmt"
)

// Represents the portions of a SQL connect string
type DbConnectSettings struct {
	DbHost, DbUser, DbPort, DbName, DbPassword string
	TLS                                        bool
}

// Formats settings into a database/sql connection string
func (c *DbConnectSettings) String() string {

	str := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s",
		c.DbHost, c.DbPort, c.DbName, c.DbUser,
	)

	if c.DbPassword != "" {
		str += fmt.Sprintf(" password=%s", c.DbPassword)
	}

	if !c.TLS {
		str += " sslmode=disable"
	}

	return str
}

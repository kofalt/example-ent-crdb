package cmd

import (
	"flag"
	"fmt"

	"github.com/kofalt/example-ent-crdb/server"
	. "github.com/kofalt/example-ent-crdb/util"
)

func RunServer(args []string, BuildHash, BuildDate string) {
	var jsonFormat bool
	var verbosity int

	var flagSet *flag.FlagSet = flag.NewFlagSet("gen", flag.ExitOnError)

	flagSet.BoolVar(&jsonFormat, "j", false, "Log output in JSON")
	flagSet.IntVar(&verbosity, "v", 0, "Log verbosity (-1 to 5)")

	Fatal(flagSet.Parse(args))

	if verbosity < -1 || verbosity > 5 {
		Fatal(fmt.Errorf("Verbosity must be from -1 to 5"))
	}

	c := &server.Config{
		LogJSON:      jsonFormat,
		LogVerbosity: verbosity,

		HttpPort:  "9000",
		BuildHash: BuildHash,
		BuildDate: BuildDate,
	}

	c.DbHost = "localhost"
	c.DbPort = "26257"
	c.DbName = "defaultdb"
	c.DbUser = "root"
	c.DbPassword = ""
	c.TLS = false

	s, err := server.Bootstrap(c)
	Fatal(err)
	s.Serve()
}

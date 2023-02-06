package main

import (
	"os"

	"github.com/kofalt/example-ent-crdb/cmd"
)

var BuildHash = "unknown"
var BuildDate = "unknown"

func main() {
	cmd.RunServer(os.Args[1:], BuildHash, BuildDate)
}

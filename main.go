package main

import (
	"flag"

	"github.com/kcmerrill/crush/core"
)

var (
	httpPort  = flag.String("web", "80", "web port")
	statsPort = flag.String("stats", "8080", "stats port")
	version   = "dev"
	commit    = "n/a"
)

func main() {
	flag.Parse()
	crush := core.CreateQ()
	go crush.Web(*httpPort)
	crush.Stat.Web(*statsPort, "")
}

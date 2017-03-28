package main

import (
	"flag"

	"github.com/kcmerrill/crush/crush"
)

var (
	port    = flag.String("port", "8080", "webserver port")
	version = "dev"
	commit  = "n/a"
)

func main() {
	flag.Parse()
	crush.CreateQ().Web(*port)
}

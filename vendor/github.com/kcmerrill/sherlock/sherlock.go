package main

import (
	"flag"

	"github.com/kcmerrill/sherlock/core"
)

var (
	udp  string
	web  string
	auth string
)

func main() {
	flag.StringVar(&udp, "udp-port", "8081", "The port in which to handle UDP requests")
	flag.StringVar(&web, "web-port", "80", "The port in which to handle WEB requests")
	flag.StringVar(&auth, "auth", "", "The authentication token used to access sherlock. No auth if left blank.")
	flag.Parse()
	s := core.New(100)
	go s.UDP(udp)
	s.Web(web, auth)
}

package main

import (
	"flag"
)

var (
	q    *Q
	port = flag.String("port", "8080", "webserver port")
)

func main() {
	flag.Parse()
	q = CreateQ()
	Web(*port)
}

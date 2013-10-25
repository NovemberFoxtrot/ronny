package main

import (
	"flag"
	"yserver"
)

func main() {
	var thePort = flag.String("port", "8080", "port to launch server on")

	flag.Parse()

	yserver.New(*thePort)
}

package main

import (
	"flag"
	"yserver"
)

func main() {
	var thePort = flag.String("port", "8080", "serve port")

	flag.Parse()

	yserver.New(*thePort)
}

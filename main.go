package main

import (
	"flag"

	"github.com/shirakiyo/FamimaGacha/internal/server"
)

var (
	port = flag.String("port", ":8080", "port to listen")
)

func main() {
	flag.Parse()
	server.Serve(*port)
}

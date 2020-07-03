package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/shirakiyo/FamimaGacha/internal/server"
)

var (
	port = flag.String("port", ":8080", "port to listen")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	flag.Parse()
	server.Serve(*port)
}

package main

import (
	"flag"
	"github.com/pythias/c10m/golang-bench/benchmarks"
)

func main() {
	host := flag.String("h", "127.0.0.1", "Server host")
	begins := flag.Int("b", 9000, "Port begins")
	ends := flag.Int("e", 9010, "Port begins")
	connections := flag.Int("n", 10000, "Number of clients to connect")
	concurrency := flag.Int("c", 10, "Number of multiple connections to make at a time")
	message := flag.String("m", "Echo Message!", "Message to echo")
	flag.Parse()

	benchmarks.StartEcho(*host, *begins, *ends, *connections, *concurrency, *message)
}



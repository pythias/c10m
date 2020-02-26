package main

import (
	"flag"
	"fmt"
	"github.com/pythias/c10m/golang-bench/benchers"
	"github.com/pythias/c10m/golang-bench/models"
	"github.com/pythias/c10m/golang-bench/utils"
	"time"
)

func main() {
	host := flag.String("h", "127.0.0.1", "Server host")
	begins := flag.Int("b", 100, "Port begins")
	connections := flag.Int("n", 10000, "Number of clients to connect")
	concurrency := flag.Int("c", 100, "Number of multiple connections to make at a time")
	message := flag.String("m", "Echo Message!", "Message to echo")
	flag.Parse()

	header := fmt.Sprintf("Bench %d connections @%s:%d-%d \n%d connections per port\nMessage is '%s'",
		*connections * *concurrency, *host, *begins, *begins + *concurrency, *connections, *message)

	counterChannel := make(chan models.Counter)

	for i := 0; i < *concurrency; i++ {
		port := *begins + i
		go func() {
			address := fmt.Sprintf("%s:%d", *host, port)
			benchers.StartEcho(address, *connections, *message, counterChannel)
		}()
	}

	results := make(map[string]models.Counter, *concurrency)
	resultTick := time.Tick(time.Second * 1)

	for {
		select {
		case counter := <-counterChannel:
			results[counter.Server] = counter
		case <-resultTick:
			utils.OutputResults(header, results)
		}
	}
}
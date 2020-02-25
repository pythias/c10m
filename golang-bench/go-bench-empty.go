package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"./utils"
)

func main() {
	timeStarted := time.Now()
	serverAddress := flag.String("s", "127.0.0.1:9003", "Server address")
	connectionNumber := flag.Int("c", 2000, "Connection number")
	flag.Parse()

	var completeNumbers, failedNumbers int64
	connectionTimes := sync.Map{}
	wg := new(sync.WaitGroup)

	for i := 0; i < *connectionNumber; i++ {
		wg.Add(1)

		go func() {
			start := time.Now()
			if _, err := net.DialTimeout("tcp", *serverAddress, time.Minute*99999); err == nil {
				elapsed := time.Since(start).Milliseconds()
				time.Sleep(time.Second * 10)
				atomic.AddInt64(&completeNumbers, 1)
				connectionTimes.Store(completeNumbers, elapsed)
			} else {
				atomic.AddInt64(&failedNumbers, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	connectionPerSecond, timePerConnection, connectionStats := utils.GetResult(connectionTimes)

	fmt.Printf("Running %0.2fs test @ %s\n", time.Since(timeStarted).Seconds(), *serverAddress)
	fmt.Println(" ", *connectionNumber, " connections")
	fmt.Printf("Complete connections: %d (%0.2f%%)\n", completeNumbers, 100*float64(completeNumbers)/float64(completeNumbers+failedNumbers))
	fmt.Println("Failed connections:", failedNumbers)
	fmt.Printf("Connections per second: %0.2f\n", connectionPerSecond)
	fmt.Printf("Time per connection: %0.2f [ms]\n", timePerConnection)
	fmt.Println("Connection Times (ms)")
	fmt.Println("             min    mean [+/-sd]  median     max")
	fmt.Println("Connect ", connectionStats)
}

package benchers

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pythias/c10m/golang-bench/utils"
)

func StartNormal(serverAddress string, connectionNumber int) {
	var completeNumbers, failedNumbers int64
	connectionTimes := sync.Map{}
	wg := new(sync.WaitGroup)
	timeStarted := time.Now()

	for i := 0; i < connectionNumber; i++ {
		wg.Add(1)

		go func() {
			start := time.Now()
			if _, err := net.DialTimeout("tcp", serverAddress, time.Minute*99999); err == nil {
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

	fmt.Printf("Running %0.2fs test @ %s\n", time.Since(timeStarted).Seconds(), serverAddress)
	fmt.Println(" ", connectionNumber, " connections")
	fmt.Printf("Complete connections: %d (%0.2f%%)\n", completeNumbers, 100*float64(completeNumbers)/float64(completeNumbers+failedNumbers))
	fmt.Println("Failed connections:", failedNumbers)
	fmt.Printf("Connections per second: %0.2f\n", connectionPerSecond)
	fmt.Printf("Time per connection: %0.2f [ms]\n", timePerConnection)
	fmt.Println("Connection Times (ms)")
	fmt.Println("             min    mean [+/-sd]  median     max")
	fmt.Println("Connect ", connectionStats)
}

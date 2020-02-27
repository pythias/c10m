package benchmarks

import (
	"fmt"
	"github.com/pythias/c10m/golang-bench/models"
	"github.com/pythias/c10m/golang-bench/utils"
	"net"
	"time"
)

func StartEcho(host string, begins int, ends int, numbers int, concurrency int, message string) {
	header := fmt.Sprintf("Benchmark v1.0\n%d connections @%s:%d-%d \nConcurrency: %d\nEcho message: %s",
		numbers, host, begins, ends, concurrency, message)

	counterChannel := make(chan models.Counter)
	connectionsPerChannel := numbers / (ends - begins + 1)

	for i := 0; i < concurrency; i++ {
		go start(i, host, begins, ends, connectionsPerChannel, message, counterChannel)
	}

	results := make(map[int]models.Counter)
	resultTick := time.Tick(time.Second * 1)

	for {
		select {
		case counter := <-counterChannel:
			results[counter.Id] = counter
		case <-resultTick:
			utils.OutputResults(header, results)
		}
	}
}

func start(id int, host string, begins int, ends int, numbers int, message string, channel chan models.Counter) {
	var connected, failed, closed, sent, received int64
	var connectionTimes []float64
	var connections []net.Conn
	for i := 0; i < numbers; i++ {
		for port := begins; port <= ends ; port++ {
			server := fmt.Sprintf("%s:%d", host, port)
			start := time.Now()
			if conn, err := net.DialTimeout("tcp", server, time.Second * 10); err == nil {
				elapsed := time.Since(start).Milliseconds()
				connected++
				connectionTimes = append(connectionTimes, float64((elapsed)))
				connections = append(connections, conn)
			} else {
				failed++
			}
		}
	}

	counter := models.Counter{
		Id:    id,
		Connected: connected,
		Failed:    failed,
	}
	stats, err := utils.CalculateStats(connectionTimes)
	if err == nil {
		counter.Stats = stats
	}
	channel <- counter

	var messageBytes = []byte(message)
	var messageLength = len(message)

	for {
		for _, conn := range connections {
			_, err1 := conn.Write(messageBytes)
			if err1 == nil {
				break
			} else {
				sent++
			}

			recv := make([]byte, messageLength)
			_, err2 := conn.Read(recv)
			if err2 == nil {
				break
			} else {
				received++
			}
		}

		counter.Sent = sent
		counter.Received = received
		counter.Closed = closed
		channel <- counter

		time.Sleep(time.Second * 60)
	}
}

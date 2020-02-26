package benchers

import (
	"github.com/pythias/c10m/golang-bench/models"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/pythias/c10m/golang-bench/utils"
)

func StartEcho(server string, connections int, message string, channel chan models.Counter) {
	var messageBytes = []byte(message)
	var messageLength = len(message)
	var connected, failed, closed, sent, received int64
	connectionTimes := sync.Map{}

	for i := 0; i < connections; i++ {
		go func() {
			start := time.Now()
			if conn, err := net.DialTimeout("tcp", server, time.Minute*99999); err == nil {
				elapsed := time.Since(start).Milliseconds()
				atomic.AddInt64(&connected, 1)
				connectionTimes.Store(connected, elapsed)

				for {
					time.Sleep(time.Second * 60)
					_, err1 := conn.Write(messageBytes)
					if err1 == nil {
						break
					} else {
						atomic.AddInt64(&sent, 1)
					}

					recv := make([]byte, messageLength)
					_, err2 := conn.Read(recv)
					if err2 == nil {
						break
					} else {
						atomic.AddInt64(&received, 1)
					}
				}
				_ = conn.Close()
				atomic.AddInt64(&closed, 1)
			} else {
				atomic.AddInt64(&failed, 1)
			}
		}()
	}

	resultTick := time.Tick(time.Second * 10)
	for {
		select {
		case <-resultTick:
			counter := models.Counter{
				Server:    server,
				Connected: connected,
				Failed:    failed,
				Closed:    closed,
				Sent:      sent,
				Received:  received,
			}
			stats, err := utils.GetResult(connectionTimes)
			if err == nil {
				counter.Stats = stats
			}

			channel <- counter
		}
	}
}


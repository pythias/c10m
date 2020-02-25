package main

import (
	"flag"

	"./benchers"
)

const ServerTypeNormal = 0
const ServerTypeEcho = 1

func main() {
	serverAddress := flag.String("s", "127.0.0.1:9003", "Server address")
	connectionNumber := flag.Int("c", 2000, "Connection number")
	serverType := flag.Int("t", ServerTypeNormal, "Server type, normal:0, echo:1")
	echoMessage := flag.String("m", "Echo Message!", "Echo message")
	flag.Parse()

	switch *serverType {
	case ServerTypeEcho:
		benchers.StartEcho(*serverAddress, *connectionNumber, *echoMessage)
	case ServerTypeNormal:
		benchers.StartNormal(*serverAddress, *connectionNumber)
	}
}

package benchers

import (
	"fmt"
)

func StartEcho(serverAddress string, connectionNumber int, message string) {
	fmt.Printf("%s, %d, %s\n", serverAddress, connectionNumber, message)
}

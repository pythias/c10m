package main

import (
	"testing"

	"github.com/pythias/c10m/golang-bench/benchers"
)

func TestEchoBench(t *testing.T) {
	benchers.StartEcho("127.0.0.1:100", 10, "Echo Message!")
}

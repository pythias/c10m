package main

import (
	"testing"

	"github.com/pythias/c10m/golang-bench/benchers"
)

func TestEchoBencher(t *testing.T) {
	benchers.StartEcho("127.0.0.1:9003", 10, "Echo Message!")
}

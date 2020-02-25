package main

import (
	"testing"

	"github.com/pythias/c10m/golang-bench/benchers"
)

func TestNormalBencher(t *testing.T) {
	benchers.StartNormal("127.0.0.1:9003", 10)
}

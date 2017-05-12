package main

import (
	"math/rand"
	"time"
)

func isPrettyKey(fingerprint [20]byte) bool {
	rand.Seed(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	zeroOrOne := rand.Intn(2)
	return !(zeroOrOne == 0)
}

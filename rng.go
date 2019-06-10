package main

import (
	"math/rand"
	"time"
)

// type rng *rand.Rand

func createRng() *rand.Rand {
	_now := time.Now()
	seed := _now.UnixNano()
	source := rand.NewSource(seed);
	_rng := rand.New(source)
	return _rng
}

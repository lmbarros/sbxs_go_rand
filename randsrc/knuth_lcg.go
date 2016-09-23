package randsrc

import "math/rand"

// knuthLCGSource is the internal state of a Linear Congruential random number
// generator.
type knuthLCGSource struct {
	state uint64
}

// NewKnuthLCG creates an unitialized rand.Source based on a Linear Congruential
// Generator (LCG) using constants recommended by Don Knuth.
func NewKnuthLCG() rand.Source {
	var rng knuthLCGSource
	return &rng
}

// Seed initializes the random number generator with a given seed.
func (rng *knuthLCGSource) Seed(seed int64) {
	rng.state = uint64(seed)
}

// Int63 generates a random number between zero and math.MaxInt64.
func (rng *knuthLCGSource) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Seed64 initializes the random number generator with a given uint64 seed,
// using all its 64 bits.
func (rng *knuthLCGSource) Seed64(seed uint64) {
	rng.state = seed
}

// Uint64 returns a random number between zero and math.MaxUint64.
func (rng *knuthLCGSource) Uint64() uint64 {
	rng.state = 6364136223846793005*rng.state + 1442695040888963407
	return rng.state
}

package randsrc

import "math/rand"

// xorshiftSource is the internal state of a Xorshift number
// generator.
type xorshiftSource struct {
	state uint64
}

// NewXorshift creates an unitialized rand.Source based on George Marsaglia's
// Xorshift algorithm.
func NewXorshift() rand.Source {
	var rng xorshiftSource
	return &rng
}

// Seed initializes the random number generator with a given seed.
func (rng *xorshiftSource) Seed(seed int64) {
	rng.state = uint64(seed)
}

// Int63 generates a random number between zero and math.MaxInt64.
func (rng *xorshiftSource) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Seed64 initializes the random number generator with a given uint64 seed,
// using all its 64 bits.
func (rng *xorshiftSource) Seed64(seed uint64) {
	rng.state = seed
}

// Uint64 returns a random number between zero and math.MaxUint64.
//
// This implements the 64-bit Xorshift version of Xorshift given as an example
// (as the xor64 function) in George Marsaglia's 2003 paper "Xorshift RNGs".
func (rng *xorshiftSource) Uint64() uint64 {
	rng.state ^= rng.state << 13
	rng.state ^= rng.state >> 7
	rng.state ^= rng.state << 17
	return rng.state
}

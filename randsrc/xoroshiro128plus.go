package randsrc

import "math/rand"

// xoroshiro128PlusSource is the internal state of a Xoroshiro128+ random number
// generator.
type xoroshiro128PlusSource struct {
	state0 uint64
	state1 uint64
}

// NewXoroshiro128Plus creates an unitialized rand.Source based on the
// Xoroshiro128+ algorithm.
//
// This is just a Go translation of the public domain C code by David Blackman
// and Sebastiano Vigna, available at
// http://xoroshiro.di.unimi.it/xoroshiro128plus.c.
func NewXoroshiro128Plus() rand.Source {
	var rng xoroshiro128PlusSource
	return &rng
}

// Seed initializes the random number generator with a given seed.
func (rng *xoroshiro128PlusSource) Seed(seed int64) {
	var sm splitMix64Source
	sm.Seed64(uint64(seed))

	rng.Seed128(sm.Uint64(), sm.Uint64())
}

// Int63 generates a random number between zero and math.MaxInt64.
func (rng *xoroshiro128PlusSource) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Seed128 initializes the random number generator with a pair of uint64 seeds,
// using all the 128 bits.
func (rng *xoroshiro128PlusSource) Seed128(seed0, seed1 uint64) {
	rng.state0 = seed0
	rng.state1 = seed1
}

// Uint64 returns a random number between zero and math.MaxUint64.
//
// This is the original function provided in the reference implementation, but
// it is not part of the rand.Source interface.
func (rng *xoroshiro128PlusSource) Uint64() uint64 {
	result := rng.state0 + rng.state1

	rng.state1 ^= rng.state0
	rng.state0 = rotl(rng.state0, 55) ^ rng.state1 ^ (rng.state1 << 14)
	rng.state1 = rotl(rng.state1, 36)

	return result
}

// rotl rotates x left by k bits.
func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

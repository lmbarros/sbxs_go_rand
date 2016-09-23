package randsrc

import "math/rand"

// splitMix64Source is the internal state of a SplitMix64 random number
// generator.
type splitMix64Source struct {
	state uint64
}

// NewSplitMix64 creates an unitialized rand.Source based on the SplitMix64
// algorithm.
//
// This is just a Go translation of the public domain C code by Sebastiano
// Vigna, available at http://xoroshiro.di.unimi.it/splitmix64.c.
func NewSplitMix64() rand.Source {
	var rng splitMix64Source
	return &rng
}

// Seed initializes the random number generator with a given seed.
func (rng *splitMix64Source) Seed(seed int64) {
	rng.state = uint64(seed)
}

// Int63 generates a random number between zero and math.MaxInt64.
func (rng *splitMix64Source) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Seed64 initializes the random number generator with a given uint64 seed,
// using all its 64 bits.
func (rng *splitMix64Source) Seed64(seed uint64) {
	rng.state = seed
}

// Uint64 returns a random number between zero and math.MaxUint64.
//
// This is the original function provided in the reference implementation, but
// it is not part of the rand.Source interface.
func (rng *splitMix64Source) Uint64() uint64 {
	rng.state += 0x9E3779B97F4A7C15
	z := rng.state
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	return z ^ (z >> 31)
}

package randsrc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lmbarros/sbxs_go_test/assert"
)

// Ensures that the rand.Source interface is properly implemented.
func TestXoroshiro128Plus_SourceInterface(t *testing.T) {
	var _ rand.Source = (*xoroshiro128PlusSource)(nil)
}

// Compares the generated pseudo random sequence with one obtained with the
// reference C implementation.
func TestXoroshiro128Plus_CheckRandomSequenceTwoSeeds(t *testing.T) {

	rng := &xoroshiro128PlusSource{}
	rng.Seed128(521288629, 362436069)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(14087298434708121260), rng.Uint64())
	assert.Equal(t, uint64(11055104410211319733), rng.Uint64())
	assert.Equal(t, uint64(4010580528061639590), rng.Uint64())
	assert.Equal(t, uint64(2807442956520957059), rng.Uint64())
	assert.Equal(t, uint64(12737219591460154268), rng.Uint64())
}

// As above, but using the single-number seeding option.
func TestXoroshiro128Plus_CheckRandomSequenceSingleSeed(t *testing.T) {

	rng := &xoroshiro128PlusSource{}

	rng.Seed(13579)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(5883773771842788424), rng.Uint64())
	assert.Equal(t, uint64(9989605189133216712), rng.Uint64())
	assert.Equal(t, uint64(7811117801491353185), rng.Uint64())
	assert.Equal(t, uint64(4706773927854861091), rng.Uint64())
	assert.Equal(t, uint64(17379387784680010921), rng.Uint64())
}

// Checks if Int63 really returns only nonnegative numbers.
func TestXoroshiro128Plus_CheckInt63(t *testing.T) {
	seed := time.Now().Unix()
	rng := NewXoroshiro128Plus()
	rng.Seed(seed)

	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Int63() >= 0,
			"Got a negative value from Int63. Seed: %v. Iteration: %v.",
			seed, i)
	}
}

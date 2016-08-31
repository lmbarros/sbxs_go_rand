package randsrc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lmbarros/sbxs_go_test/assert"
)

// Ensures that the rand.Source interface is properly implemented.
func TestSplitMix64_SourceInterface(t *testing.T) {
	var _ rand.Source = (*splitMix64Source)(nil)
}

// Compares the generated pseudo random sequence with one obtained with the
// reference C implementation.
func TestSplitMix64_CheckRandomSequence112233(t *testing.T) {

	rng := &splitMix64Source{}

	rng.Seed64(112233)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(17521931875799178301), rng.Uint64())
	assert.Equal(t, uint64(14383456051838709316), rng.Uint64())
	assert.Equal(t, uint64(2828063782273925862), rng.Uint64())
	assert.Equal(t, uint64(11956729851173419597), rng.Uint64())
	assert.Equal(t, uint64(15283991741788426049), rng.Uint64())
}

// Compares the generated pseudo random sequence with one obtained with the
// reference C implementation.
func TestSplitMix64_CheckRandomSequence97531(t *testing.T) {

	rng := &splitMix64Source{}

	rng.Seed64(97531)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(14309215551682402496), rng.Uint64())
	assert.Equal(t, uint64(13331760577475424929), rng.Uint64())
	assert.Equal(t, uint64(6476874722340802351), rng.Uint64())
	assert.Equal(t, uint64(13089128203191675115), rng.Uint64())
	assert.Equal(t, uint64(7925169073796594034), rng.Uint64())
}

// Checks if Int63 really returns only nonnegative numbers.
func TestSplitMix64_CheckInt63(t *testing.T) {
	seed := time.Now().Unix()
	rng := NewSplitMix64()
	rng.Seed(seed)

	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Int63() >= 0,
			"Got a negative value from Int63. Seed: %v. Iteration: %v.",
			seed, i)
	}
}

// Just because I am new to this Go thing, use the source with some of the
// standard rand functions.
func TestSplitMix64_CheckSource(t *testing.T) {
	rng := rand.New(NewSplitMix64())
	seed := time.Now().Unix()

	rng.Seed(seed)
	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Float32() < 1.0,
			"Got a negative value from Float32. Seed: %v. Iteration: %v.",
			seed, i)
	}

	rng.Seed(seed)
	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Intn(10) < 10,
			"Got an incorrect value from Intn. Seed: %v. Iteration: %v.",
			seed, i)
	}
}

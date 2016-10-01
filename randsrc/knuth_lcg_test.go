package randsrc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lmbarros/sbxs_go_test/assert"
)

// Ensures that the rand.Source interface is properly implemented.
func TestKnuthLCG_SourceInterface(t *testing.T) {
	var _ rand.Source = (*knuthLCGSource)(nil)
}

// Compares the generated pseudo random sequence with one obtained with a C
// implementation I found somewhere.
func TestKnuthLCG_CheckRandomSequence(t *testing.T) {

	rng := &knuthLCGSource{}
	rng.Seed64(112233)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(14373087394460283212), rng.Uint64())
	assert.Equal(t, uint64(17919820627726294955), rng.Uint64())
	assert.Equal(t, uint64(2374928239783487326), rng.Uint64())
	assert.Equal(t, uint64(1355612529541287637), rng.Uint64())
	assert.Equal(t, uint64(14300717203320031168), rng.Uint64())

	// Again, with a different seed
	rng.Seed64(97531)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(2475730772265627958), rng.Uint64())
	assert.Equal(t, uint64(16585264879776313805), rng.Uint64())
	assert.Equal(t, uint64(10829226592693777752), rng.Uint64())
	assert.Equal(t, uint64(4337668493995821511), rng.Uint64())
	assert.Equal(t, uint64(13619766792622649930), rng.Uint64())
}

// Checks if Int63 really returns only nonnegative numbers.
func TestKnuthLCG_CheckInt63(t *testing.T) {
	seed := time.Now().Unix()
	rng := NewKnuthLCG()
	rng.Seed(seed)

	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Int63() >= 0,
			"Got a negative value from Int63. Seed: %v. Iteration: %v.",
			seed, i)
	}
}

// Benchmarks the LCG algorithm.
func BenchmarkKnuthLCG(b *testing.B) {
	rng := NewKnuthLCG()
	rng.Seed(12345)

	for i := 0; i < b.N; i++ {
		rng.Int63()
	}
}

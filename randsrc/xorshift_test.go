package randsrc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lmbarros/sbxs_go_test/assert"
)

// Ensures that the rand.Source interface is properly implemented.
func TestXorshift_SourceInterface(t *testing.T) {
	var _ rand.Source = (*xorshiftSource)(nil)
}

// Compares the generated pseudo random sequence with the reference C
// implementation (from Marsaglia's 2003 paper "Xorshift RNGs").
func TestXorshift_CheckRandomSequence(t *testing.T) {

	rng := &xorshiftSource{}
	rng.Seed64(112233)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(15708644121595424226), rng.Uint64())
	assert.Equal(t, uint64(5588328912753616697), rng.Uint64())
	assert.Equal(t, uint64(1369163389825252279), rng.Uint64())
	assert.Equal(t, uint64(5877863171656316800), rng.Uint64())
	assert.Equal(t, uint64(4133985236543262319), rng.Uint64())

	// Again, with a different seed
	rng.Seed64(97531)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(7406447624127097846), rng.Uint64())
	assert.Equal(t, uint64(3225819606517077281), rng.Uint64())
	assert.Equal(t, uint64(6624085062535109835), rng.Uint64())
	assert.Equal(t, uint64(13324435952915893938), rng.Uint64())
	assert.Equal(t, uint64(15108874822966239151), rng.Uint64())
}

// Checks if Int63 really returns only nonnegative numbers.
func TestXorshift_CheckInt63(t *testing.T) {
	seed := time.Now().Unix()
	rng := NewXorshift()
	rng.Seed(seed)

	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Int63() >= 0,
			"Got a negative value from Int63. Seed: %v. Iteration: %v.",
			seed, i)
	}
}

// Benchmarks the Xorshift algorithm.
func BenchmarkXorshift(b *testing.B) {
	rng := NewXorshift()
	rng.Seed(12345)

	for i := 0; i < b.N; i++ {
		rng.Int63()
	}
}

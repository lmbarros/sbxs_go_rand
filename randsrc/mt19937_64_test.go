package randsrc

import (
	"math/rand"
	"testing"
	"time"

	"github.com/lmbarros/sbxs_go_test/assert"
)

// Ensures that the rand.Source interface is properly implemented.
func TestMT19937_64_SourceInterface(t *testing.T) {
	var _ rand.Source = (*mt19937_64Source)(nil)
}

// Compares the generated pseudo random sequence with one obtained with the
// reference C implementation (found at
// http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/VERSIONS/C-LANG/mt19937-64.c).
func TestMT19937_64_CheckRandomSequence(t *testing.T) {

	rng := &mt19937_64Source{}
	rng.Seed64(112233)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(18385898271542208877), rng.Uint64())
	assert.Equal(t, uint64(10388062775371028914), rng.Uint64())
	assert.Equal(t, uint64(6652856830127087739), rng.Uint64())
	assert.Equal(t, uint64(1890554855370499999), rng.Uint64())
	assert.Equal(t, uint64(14650261722016513395), rng.Uint64())

	// Again, with a different seed
	rng.Seed64(97531)

	for i := 0; i < 1000; i++ {
		rng.Uint64()
	}

	assert.Equal(t, uint64(6788201074649841524), rng.Uint64())
	assert.Equal(t, uint64(11921711492735785818), rng.Uint64())
	assert.Equal(t, uint64(9983243477236236108), rng.Uint64())
	assert.Equal(t, uint64(5762306495090730538), rng.Uint64())
	assert.Equal(t, uint64(636925918925826639), rng.Uint64())
}

// Checks if Int63 really returns only nonnegative numbers.
func TestMT19937_64_CheckInt63(t *testing.T) {
	seed := time.Now().Unix()
	rng := NewMT19937_64()
	rng.Seed(seed)

	for i := 0; i < 10000; i++ {
		assert.True(t, rng.Int63() >= 0,
			"Got a negative value from Int63. Seed: %v. Iteration: %v.",
			seed, i)
	}
}

// Benchmarks the MT19937-64 algorithm.
func BenchmarkMT19937_64(b *testing.B) {
	rng := NewMT19937_64()
	rng.Seed(12345)

	for i := 0; i < b.N; i++ {
		rng.Int63()
	}
}

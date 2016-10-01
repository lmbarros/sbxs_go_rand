// Implementation based on the pseudo-code from Wikipedia
// (https://en.wikipedia.org/wiki/Mersenne_Twister).

package randsrc

import "math/rand"

const (
	w = 64
	n = 312
	m = 156
	r = 31

	a = 0xB5026F5AA96619E9
	u = 29
	d = 0x5555555555555555
	s = 17
	b = 0x71D67FFFEDA60000
	t = 37
	c = 0xFFF7EEE000000000
	l = 43

	f = 6364136223846793005

	lowerMask = uint64((1 << r) - 1)
	upperMask = ^lowerMask
)

// mt19937_64Source is the internal state of a MT19937-64 random number generator.
type mt19937_64Source struct {
	mt    [n]uint64
	index uint
}

// NewMT19937_64 creates an unitialized rand.Source based on the MT19937-64
// version of the Mersenne Twister algorithm.
func NewMT19937_64() rand.Source {
	var rng mt19937_64Source
	rng.index = n + 1
	return &rng
}

// Seed initializes the random number generator with a given seed.
func (rng *mt19937_64Source) Seed(seed int64) {
	rng.Seed64(uint64(seed))
}

// Int63 generates a random number between zero and math.MaxInt64.
func (rng *mt19937_64Source) Int63() int64 {
	return int64(rng.Uint64() >> 1)
}

// Seed64 initializes the random number generator with a given uint64 seed.
func (rng *mt19937_64Source) Seed64(seed uint64) {
	rng.index = n
	rng.mt[0] = seed
	for i := uint64(1); i < n; i++ {
		rng.mt[i] = f*(rng.mt[i-1]^(rng.mt[i-1]>>(w-2))) + i
	}
}

// Uint64 returns a random number between zero and math.MaxUint64.
func (rng *mt19937_64Source) Uint64() uint64 {
	if rng.index >= n {
		if rng.index > n {
			rng.Seed64(5489)
		}
		rng.twist()
	}

	y := rng.mt[rng.index]
	y = y ^ ((y >> u) & d)
	y = y ^ ((y << s) & b)
	y = y ^ ((y << t) & c)
	y = y ^ (y >> l)

	rng.index++
	return y
}

// twist twists the Mersennes, I suppose.
func (rng *mt19937_64Source) twist() {
	for i := 0; i < n; i++ {
		x := (rng.mt[i] & upperMask) + (rng.mt[(i+1)%n] & lowerMask)
		xA := x >> 1
		if (x % 2) != 0 { // lowest bit of x is 1
			xA = xA ^ a
		}
		rng.mt[i] = rng.mt[(i+m)%n] ^ xA
	}
	rng.index = 0
}

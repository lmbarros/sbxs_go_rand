package randsrc

import (
	"fmt"
	"unsafe"
)

// Not a real example, but rather a simple tool to check each algorithm's state
// size.
func Example_sizeof() {
	fmt.Printf("KnuthLCG: %v\n", unsafe.Sizeof(knuthLCGSource{}))
	fmt.Printf("MT19937-64: %v\n", unsafe.Sizeof(mt19937_64Source{}))
	fmt.Printf("SplitMix64: %v\n", unsafe.Sizeof(splitMix64Source{}))
	fmt.Printf("Xoroshiro128+: %v\n", unsafe.Sizeof(xoroshiro128PlusSource{}))
	fmt.Printf("Xorshift: %v\n", unsafe.Sizeof(xorshiftSource{}))

	// Output:
	// KnuthLCG: 8
	// MT19937-64: 2504
	// SplitMix64: 8
	// Xoroshiro128+: 16
	// Xorshift: 8
}

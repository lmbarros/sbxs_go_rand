package randutil

import (
	"os"
	"time"
)

// GoodSeed returns a good value for seeding a random number generator.
func GoodSeed() int64 {
	// Use time-based values for the higher 48 bits
	high := uint64(time.Now().Unix() << 32)
	mid := uint64((time.Now().UnixNano() & 0xFFFF0000) << 16)

	// Including the PID in the seed is a good idea because it warrants us
	// distinct seeds if called in two processes created at the same time.
	// This is unlikely to happen given the use of UnixNano, but better safe
	// than sorry. (Failing to use the PID in the seed caused me bugs a long
	// time ago, but at that time I was using the time only with seconds
	// resolution.)
	low := uint64(os.Getpid() & 0x0000FFFF)

	return int64(high | mid | low)
}

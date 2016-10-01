# StackedBoxes' Random library in Go

[![GoDoc](https://godoc.org/github.com/lmbarros/sbxs_go_rand?status.svg)](https://godoc.org/github.com/lmbarros/sbxs_go_rand) [![Go Report Card](https://goreportcard.com/badge/github.com/lmbarros/sbxs_go_rand)](https://goreportcard.com/report/github.com/lmbarros/sbxs_go_rand) ![License](https://img.shields.io/github/license/lmbarros/sbxs_go_rand.svg)

Random number utilities in Go.

Package `randsrc` provides (pseudo) random number sources based on the following
algorithms:

- **Xoroshiro128+**: This would make a good first choice of algorithm for most
  applications which don't depend critically on random numbers and will not get
  people killed or broken if something fails. It is quite fast, uses only 16
  bytes of state and has a decent (if not astronomical) period of 2^128 - 1. It
  also passes lots of hard randomness tests. If you are curious (like I was),
  "xoroshiro" stands for "XOr/ROtate/SHIft/ROtate".

- **MT19937-64**: The 64-bit version of [Mersenne
  Twister](http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt64.html) with
  multiplier and increment values as suggested by Don Knuth himself (or so says
  the Wikipedia). Notice that even though this has "Knuth" in the name, this is
  still an LCG, and therefore you probably don't want to use it.

- **SplitMix64**: Not bad generator at all, but Xoroshiro128+ wins in almost
  every aspect. SplitMix64 wins in memory usage for storing the state, but we
  are talking about 8 versus 16 bytes). It also seems to be a bit faster than
  Xoroshiro128+, at least in these Go implementations. There another good use
  for this algorithm: suppose you want to seed an algorithm that accepts very
  large seeds, but all you have is a humble int64. Just use your int64 to
  initialize a SplitMix64 and use it to generate as much seed data as you need.

- **Knuth LCG**: A [Linear Congruential
  Generator](https://en.wikipedia.org/wiki/Linear_congruential_generator) (LCG)
  with multiplier and increment values as suggested by Don
  Knuth himself (or so says the Wikipedia). Notice that even though this has
  "Knuth" in the name, this is still an LCG, and therefore you probably don't
  want to use it.


Package `randutil` provides assorted utilities for working with random numbers.
Currently, it contains only `GoodSeed`, a function providing a good value to use
as a random seed.

## License

All code here is under the MIT License.

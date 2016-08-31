# sbxs_go_rand

Random number utilities in Go.

Package randsrc provides (pseudo) random number sources based on the following
algorithms:

- **SplitMix64**: Not bad generator at all, but Xoroshiro128+ wins in almost
  every aspect. SplitMix64 wins in memory usage for storing the state, but we
  are talking about 8 versus 16 bytes). There is one good use for this
  algorithm: suppose you want to seed an algorithm that accepts very large
  seeds, but all you have is a humble int64. Just use your int64 to initialize
  a SplitMix64 and use it to generate as much seed data as you need.

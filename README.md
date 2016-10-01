# StackedBoxes' Random library in Go

[![GoDoc](https://godoc.org/github.com/lmbarros/sbxs_go_rand?status.svg)](https://godoc.org/github.com/lmbarros/sbxs_go_rand) [![Go Report Card](https://goreportcard.com/badge/github.com/lmbarros/sbxs_go_rand)](https://goreportcard.com/report/github.com/lmbarros/sbxs_go_rand) ![License](https://img.shields.io/github/license/lmbarros/sbxs_go_rand.svg)

Random number utilities in Go.

Package `randsrc` provides (pseudo) random number sources based on the following
algorithms:

<table>
    <tr>
        <th>Name</th>
        <th>Speed (ns/op)</th>
        <th>Period</th>
        <th>Size (bytes)</th>
        <th>Notes</th>
        <th>References</th>
    </tr>

    <tr>
        <td>Knuth LCG</td>
        <td>2.71</td>
        <td>No more than 2<sup>64</sup></td>
        <td>8</td>
        <td>An LCG with constants suggested by Don Knuth (or so says Wikipedia). Low quality, despite having "Knuth" in the name.</td>
        <td><a href=https://en.wikipedia.org/wiki/Linear_congruential_generator>WP</a></td>
    </tr>

    <tr>
        <td>MT19937-64</td>
        <td>15.2</td>
        <td>2<sup>19937</sup>−1</td>
        <td>2504</td>
        <td>Good quality. Very popular. Ridiculously long period. Large state.</td>
        <td><a href="https://en.wikipedia.org/wiki/Mersenne_Twister">WP</a>, <a href="http://www.math.sci.hiroshima-u.ac.jp/~m-mat/MT/emt.html">Official</a></td>
    </tr>

    <tr>
        <td>SplitMix64</td>
        <td>3.28</td>
        <td>2<sup>64</sup></td>
        <td>8</td>
        <td>Not bad. Good choice for seeding other RNGs from a single 64-bit number.</td>
        <td><a href="http://xoroshiro.di.unimi.it">Link</a></td>
    </tr>

    <tr>
        <td>Xoroshiro128+</td>
        <td>4.43</td>
        <td>2<sup>128</sup>−1</td>
        <td>16</td>
        <td>A very good one. "Xoroshiro" stands for "XOr/ROtate/SHIft/ROtate".</td>
        <td><a href="http://xoroshiro.di.unimi.it/">Official</a></td>
    </tr>

    <tr>
        <td>Xorshift</td>
        <td>3.49</td>
        <td>2<sup>64</sup>−1</td>
        <td>8</td>
        <td>A good one. Marsaglia was the man.</td>
        <td><a href="https://en.wikipedia.org/wiki/Xorshift">WP</a>, <a href="www.jstatsoft.org/v08/i14/paper">Paper</a></td>
    </tr>
</table>


Package `randutil` provides assorted utilities for working with random numbers.
Currently, it contains only `GoodSeed`, a function providing a good value to use
as a random seed.

## License

All code here is under the MIT License.

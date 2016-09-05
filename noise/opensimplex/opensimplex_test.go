// Note by Leandro Motta Barros: The nice tests for OpenSimplex Noise were
// originally written by Owen Raccuglia. They kind of go in the same vein as
// the tests I did in my D (dlang) OpenSimples Nose implementation (see
// https://github.com/lmbarros/sbxs_dlang/blob/master/src/sbxs/noise/open_simplex_noise.d),
// but the credits of this code go to Owen (who clearly knows Go much better
// than I do, by the way), not to me, who just added a few simple benchmarks.

// Tests for OpenSimplex noise, based on the output of
// the Java implementation.
//
// All reference samples were rendered with the default seed (0). Each version
// of the noise function (2D, 3D and 4D) was run to output 2D samples slicing
// across two of the function's axes. There is one 2D slice, three 3D slices
// and 6 4D slices; the 3D slices each pin one axis to the value 3.8; 4D slices
// pin one axis (the first in the filename) to 3.8 and the second to 2.7. These
// values were chosen arbitrarily.
//
// Each sample is a 512x512 greyscale PNG; each pixel is 1/24 wide in the
// OpenSimplex's space -- i.e. pixel (24, 24) in the 2D noise sample was
// computed by evaluating the 2D noise at (1.0, 1.0) and converting from a [-1,
// +1] scale to [0, +1].
//
package opensimplex

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"os"
	"path"
	"testing"

	"github.com/lmbarros/sbxs_go_test/test/assert"
)

func loadSamples() <-chan []float64 {
	c := make(chan []float64)
	go func() {
		f, err := os.Open(path.Join("test_files", "samples.json.gz"))
		if err != nil {
			panic(err.Error())
		}
		defer f.Close()

		gz, err := gzip.NewReader(f)
		if err != nil {
			panic(err.Error())
		}

		dec := json.NewDecoder(gz)
		for {
			var sample []float64
			if err := dec.Decode(&sample); err == io.EOF {
				break
			} else if err != nil {
				panic(err.Error())
			} else {
				c <- sample
			}
		}
		close(c)
	}()

	return c
}

// Compares generated noise values with values generated with the reference Java
// implementation.
func TestSamplesMatch(t *testing.T) {
	samples := loadSamples()
	n := NewWithSeed(0)

	for s := range samples {
		var expected, actual float64
		switch len(s) {
		case 3:
			expected = s[2]
			actual = n.Noise2D(s[0], s[1])
		case 4:
			expected = s[3]
			actual = n.Noise3D(s[0], s[1], s[2])
		case 5:
			expected = s[4]
			actual = n.Noise4D(s[0], s[1], s[2], s[3])
		default:
			t.Fatalf("Unexpected size sample: %d", len(s))
		}

		if expected != actual {
			t.Fatalf("Expected %v, got %v for %dD sample at %v",
				expected, actual, len(s)-1, s[:len(s)-1])
		}
	}
}

// Makes sure that the 1D noise behaves as if sampling 2D noise at y = 0.0. This
// test serves to allow me to try to optmize my "fake 1D" implementation while
// ensuring that I didn't mess things up.
func Test1DNoise(t *testing.T) {
	noise := New()

	for x := -10.0; x < 10.0; x += 0.09 {
		assert.Equal(t, noise.Noise2D(x, 0.0), noise.Noise1D(x))
	}
}

// Benchmarks 1D noise generation
func Benchmark1D(b *testing.B) {
	noise := New()

	for i := 0; i < b.N; i++ {
		noise.Noise1D(float64(i))
	}
}

// Benchmarks 2D noise generation
func Benchmark2D(b *testing.B) {
	noise := New()

	for i := 0; i < b.N; i++ {
		noise.Noise2D(float64(i), float64(i))
	}
}

// Benchmarks 3D noise generation
func Benchmark3D(b *testing.B) {
	noise := New()

	for i := 0; i < b.N; i++ {
		noise.Noise3D(float64(i), float64(i), float64(i))
	}
}

// Benchmarks 4D noise generation
func Benchmark4D(b *testing.B) {
	noise := New()

	for i := 0; i < b.N; i++ {
		noise.Noise4D(float64(i), float64(i), float64(i), float64(i))
	}
}

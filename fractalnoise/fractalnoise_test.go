package fractalnoise

import (
	"testing"

	"github.com/lmbarros/sbxs_go_noise"
	"github.com/lmbarros/sbxs_go_test/test/assert"
)

// mockedNoise is a fake noise generator, which always returns 1.0.
type mockedNoise struct{}

func (n *mockedNoise) Noise1D(x float64) float64 {
	return 1.0
}

func (n *mockedNoise) Noise2D(x, y float64) float64 {
	return 1.0
}

func (n *mockedNoise) Noise3D(x, y, z float64) float64 {
	return 1.0
}

func (n *mockedNoise) Noise4D(x, y, z, w float64) float64 {
	return 1.0
}

// Tests 1D fractal noise.
func TestFractalNoise1D(t *testing.T) {
	// 1D, default gain of 0.5
	mn := noise.Noiser1D(&mockedNoise{})
	fn := New1D(mn, Params{Layers: 2})
	noise := fn.Noise1D(0.0)

	assert.Equal(t, noise, 1.5)
}

// Tests 2D fractal noise.
func TestFractalNoise2D(t *testing.T) {
	// 2D, default gain of 0.5
	mn := noise.Noiser2D(&mockedNoise{})
	fn := New2D(mn, Params{Layers: 5})
	noise := fn.Noise2D(0.0, 0.0)

	assert.Equal(t, noise, 1.9375)
}

// Tests 3D fractal noise.
func TestFractalNoise3D(t *testing.T) {
	// 3D, gain = 0.8, default 4 layers
	mn := noise.Noiser3D(&mockedNoise{})
	fn := New3D(mn, Params{Gain: 0.8})
	noise := fn.Noise3D(0.0, 0.0, 0.0)

	epsilon := 1e-7
	assert.Close64(t, noise, 2.952, epsilon)
}

// Tests 4D fractal noise.
func TestFractalNoise4D(t *testing.T) {
	// 4D, gain = 0.1
	mn := noise.Noiser4D(&mockedNoise{})
	fn := New4D(mn, Params{Layers: 3, Gain: 0.1})
	noise := fn.Noise4D(0.0, 0.0, 0.0, 0.0)

	assert.Equal(t, noise, 1.11)
}
